// lib/config/config.dart

// API Configuration
class ApiConfig {
  static const String baseUrl = "http://127.0.0.1:8080";
  static const String apiKey = "your-api-key";

  static const String getFeeData = "/get-fee-data";
  static const String submitPayment = "/submit-payment";

  static String get baseApiUrl => baseUrl; // Returning base URL
}

// App Configuration
class AppConfig {
  static const String appName = "Swift School";
  static const String appVersion = "1.0.0";
  static const String appDescription =
      "A simple app to manage fee counters and transactions.";
}
