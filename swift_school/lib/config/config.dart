// lib/config/config.dart

// API Configuration
import 'dart:io' show Platform;
import 'package:flutter/foundation.dart' show kIsWeb, kDebugMode;

class ApiConfig {
  static String get baseUrl {
    final url = _getBaseUrl();
    if (kDebugMode) {
      print('Platform: ${_getPlatformInfo()}');
      print('Using API URL: $url');
    }
    return url;
  }

  static String _getBaseUrl() {
    if (kIsWeb) {
      return "http://localhost:8080";
    } else if (Platform.isAndroid) {
      return "http://10.0.2.2:8080";
    } else if (Platform.isMacOS) {
      return "http://127.0.0.1:8080"; // Specific for macOS desktop
    } else {
      return "http://localhost:8080"; // Other platforms
    }
  }

  static String _getPlatformInfo() {
    if (kIsWeb) return 'Web';
    if (Platform.isAndroid) return 'Android';
    if (Platform.isIOS) return 'iOS';
    if (Platform.isMacOS) return 'macOS';
    if (Platform.isWindows) return 'Windows';
    if (Platform.isLinux) return 'Linux';
    return 'Unknown';
  }

  static const String apiKey = "your-api-key";
  static const String getFeeData = "/get-fee-data";
  static const String submitPayment = "/submit-payment";

  // API Endpoints
  static const String healthCheck = "/api/health";
  static const String getData = "/api/data";
  static const String feeStructure = "/api/fees/fee-structure-by-class";
  static const String studentRegistration = "/api/student/student-registration";
}

// App Configuration
class AppConfig {
  static const String appName = "Swift School";
  static const String appVersion = "1.0.0";
  static const String appDescription =
      "A simple app to manage fee counters and transactions.";
}
