import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'package:http/io_client.dart';
import 'package:swift_school/config/config.dart';
import '../models/data.dart';
import 'package:connectivity_plus/connectivity_plus.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiException implements Exception {
  final String message;
  final int? statusCode;
  final String? details;

  ApiException(this.message, {this.statusCode, this.details});

  @override
  String toString() =>
      'ApiException: $message${statusCode != null ? ' (Status: $statusCode)' : ''}${details != null ? '\nDetails: $details' : ''}';
}

class ApiService {
  final String baseUrl;
  final http.Client _client;
  final Duration timeout;
  final int maxRetries;

  static const String _cacheKey = 'api_cache_';
  static const Duration _cacheExpiration = Duration(minutes: 5);

  /// Creates an instance of ApiService
  ApiService._({
    required this.baseUrl,
    http.Client? client,
    this.timeout = const Duration(seconds: 30),
    this.maxRetries = 3,
  }) : _client = client ?? http.Client() {
    if (kDebugMode) {
      print('Initializing ApiService with:');
      print('- Base URL: $baseUrl');
      print('- Platform: ${Platform.operatingSystem}');
      print('- Timeout: ${timeout.inSeconds} seconds');
      print('- Max retries: $maxRetries');
    }
  }

  /// Factory constructor for creating ApiService instances
  factory ApiService({
    String? baseUrl,
    http.Client? client,
    Duration? timeout,
    int? maxRetries,
  }) {
    // For desktop, try different localhost variants
    if (Platform.isMacOS) {
      if (kDebugMode) {
        print('Initializing for macOS platform');
      }

      // Create an HttpClient that forces IPv4
      final httpClient = HttpClient()
        ..connectionTimeout = const Duration(seconds: 5)
        ..badCertificateCallback =
            (cert, host, port) => true; // Allow self-signed certs for local dev

      // Create an IOClient that uses our configured HttpClient
      final ioClient = IOClient(httpClient);

      return ApiService._(
        baseUrl: baseUrl ?? 'http://127.0.0.1:8080', // Force IPv4
        client: ioClient,
        timeout: timeout ?? const Duration(seconds: 30),
        maxRetries: maxRetries ?? 3,
      );
    }

    // For other platforms, use the default URL
    return ApiService._(
      baseUrl: baseUrl ?? ApiConfig.baseUrl,
      client: client,
      timeout: timeout ?? const Duration(seconds: 30),
      maxRetries: maxRetries ?? 3,
    );
  }

  /// Tests connectivity to a URL
  static Future<bool> testConnection(String url) async {
    if (kDebugMode) {
      print('Testing connection to: $url');
    }

    final httpClient = HttpClient()
      ..connectionTimeout = const Duration(seconds: 5)
      ..badCertificateCallback = (cert, host, port) => true;

    try {
      final request = await httpClient.getUrl(Uri.parse('$url/api/health'));
      request.headers.add('Accept', 'application/json');
      final response = await request.close();
      await response.drain(); // Drain the response to free resources
      return response.statusCode == 200;
    } catch (e) {
      if (kDebugMode) {
        print('Connection test failed for $url: $e');
      }
      return false;
    } finally {
      httpClient.close();
    }
  }

  Future<bool> _checkConnectivity() async {
    var connectivityResult = await Connectivity().checkConnectivity();
    // ignore: unrelated_type_equality_checks
    return connectivityResult != ConnectivityResult.none;
  }

  Future<T> _withRetry<T>(Future<T> Function() operation) async {
    int attempts = 0;
    while (attempts < maxRetries) {
      try {
        return await operation();
      } catch (e) {
        attempts++;
        if (attempts == maxRetries) rethrow;

        if (e is SocketException || e is http.ClientException) {
          await Future.delayed(
              Duration(seconds: attempts * 2)); // Exponential backoff
          continue;
        }
        rethrow;
      }
    }
    throw ApiException('Max retry attempts reached');
  }

  Future<Map<String, String>> _getHeaders() async {
    final uri = Uri.parse(baseUrl);
    return {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
      'Connection': 'keep-alive',
      'Access-Control-Allow-Origin': '*',
      'User-Agent': 'Flutter/Desktop',
      'Host': uri.host,
      'X-Requested-With': 'XMLHttpRequest',
    };
  }

  Future<void> _cacheResponse(String endpoint, String data) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final cacheData = {
        'data': data,
        'timestamp': DateTime.now().millisecondsSinceEpoch,
      };
      await prefs.setString(_cacheKey + endpoint, json.encode(cacheData));
    } catch (e) {
      if (kDebugMode) {
        print('Cache write error: $e');
      }
    }
  }

  Future<String?> _getCachedResponse(String endpoint) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final cachedData = prefs.getString(_cacheKey + endpoint);
      if (cachedData != null) {
        final decoded = json.decode(cachedData);
        final timestamp = decoded['timestamp'] as int;
        if (DateTime.now().millisecondsSinceEpoch - timestamp <
            _cacheExpiration.inMilliseconds) {
          return decoded['data'];
        }
      }
    } catch (e) {
      if (kDebugMode) {
        print('Cache read error: $e');
      }
    }
    return null;
  }

  /// Gets the full URL for an endpoint
  String _getFullUrl(String endpoint) {
    final url =
        endpoint.startsWith('/') ? '$baseUrl$endpoint' : '$baseUrl/$endpoint';
    if (kDebugMode) {
      print('Generated URL: $url');
    }
    return url;
  }

  Future<List<Data>> fetchData({bool forceRefresh = false}) async {
    if (kDebugMode) {
      print('Fetching data with following configuration:');
      print('Base URL: $baseUrl');
      print('Timeout: ${timeout.inSeconds} seconds');
      print('Max retries: $maxRetries');
    }

    if (!await _checkConnectivity()) {
      throw ApiException(
          'No internet connection. Please check your network settings.');
    }

    try {
      // Check cache first if not forcing refresh
      if (!forceRefresh) {
        final cachedData = await _getCachedResponse(ApiConfig.getData);
        if (cachedData != null) {
          if (kDebugMode) {
            print('Returning cached data');
          }
          final List<dynamic> decodedData = json.decode(cachedData);
          return decodedData.map((json) => Data.fromJson(json)).toList();
        }
      }

      // Make API call with retry mechanism
      return await _withRetry(() async {
        final url = _getFullUrl(ApiConfig.getData);
        if (kDebugMode) {
          print('Making HTTP GET request to: $url');
        }

        final headers = await _getHeaders();
        if (kDebugMode) {
          print('Request headers: $headers');
        }

        try {
          final response = await _client
              .get(Uri.parse(url), headers: headers)
              .timeout(timeout);

          if (kDebugMode) {
            print('Response status code: ${response.statusCode}');
            print('Response headers: ${response.headers}');
            if (response.statusCode != 200) {
              print('Response body: ${response.body}');
            }
          }

          if (response.statusCode == 200) {
            // Cache the successful response
            await _cacheResponse(ApiConfig.getData, response.body);

            List<dynamic> dataJson = json.decode(response.body);
            return dataJson.map((json) => Data.fromJson(json)).toList();
          } else if (response.statusCode == 404) {
            throw ApiException(
                'Resource not found. Please check if the server is running and the endpoint is correct.',
                statusCode: response.statusCode);
          } else if (response.statusCode >= 500) {
            throw ApiException('Server error. Please try again later.',
                statusCode: response.statusCode, details: response.body);
          } else {
            throw ApiException('Failed to load data',
                statusCode: response.statusCode,
                details:
                    'Unexpected status code: ${response.statusCode}\nBody: ${response.body}');
          }
        } on SocketException catch (e) {
          if (kDebugMode) {
            print('Socket Exception:');
            print('  Message: ${e.message}');
            print('  Address: ${e.address}');
            print('  Port: ${e.port}');
            print('  OS Error: ${e.osError}');
          }
          throw ApiException(
              'Network error. Please check if the server is running on $url',
              details: 'Error: ${e.message}\nPort: ${e.port}');
        }
      });
    } on TimeoutException {
      throw ApiException('Request timed out after ${timeout.inSeconds} seconds',
          details: 'The server at $baseUrl is not responding');
    } on FormatException catch (e) {
      throw ApiException('Invalid response format',
          details: 'The server response could not be parsed: ${e.message}');
    } catch (e) {
      if (e is ApiException) rethrow;
      throw ApiException('Unexpected error', details: e.toString());
    }
  }

  void dispose() {
    _client.close();
  }
}
