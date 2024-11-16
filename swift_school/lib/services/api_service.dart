import 'dart:convert';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'package:http/http.dart';
import 'package:swift_school/config/config.dart';
import '../models/data.dart';

class ApiService {
  final String baseUrl;

  ApiService({this.baseUrl = ApiConfig.baseUrl});

  Future<List<Data>> fetchData() async {
    try {
      if (kDebugMode) {
        print('Fetching data from: $baseUrl/api/data');
      }
      final response = await http.get(Uri.parse('$baseUrl/api/data'));

      if (response.statusCode == 200) {
        List<dynamic> dataJson = json.decode(response.body);
        return dataJson.map((json) => Data.fromJson(json)).toList();
      } else {
        throw Exception('Failed to load data: ${response.statusCode}');
      }
    } catch (e) {
      if (kDebugMode) {
        print('Error occurred: $e');
      } // Print detailed error for debugging
      if (e is ClientException) {
        if (kDebugMode) {
          print("ClientException: $e");
        }
      }
      if (e is SocketException) {
        if (kDebugMode) {
          print("SocketException: $e");
        }
      }
      if (e is FormatException) {
        if (kDebugMode) {
          print("FormatException: $e");
        }
      }
      throw Exception('Failed to fetch data: $e');
    }
  }
}
