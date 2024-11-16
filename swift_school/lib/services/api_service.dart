import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/data.dart';

class ApiService {
  final String baseUrl;

  ApiService({this.baseUrl = 'http://127.0.0.1:8080'});

  Future<List<Data>> fetchData() async {
    try {
      // Log the URL being requested
      print('Fetching data from: $baseUrl/api/data');

      final response = await http.get(Uri.parse('$baseUrl/api/data'));

      // Log the status code of the response
      print('Response Status: ${response.statusCode}');

      if (response.statusCode == 200) {
        List<dynamic> dataJson = json.decode(response.body);

        // Log the response body
        print('Response Body: ${response.body}');

        return dataJson.map((json) => Data.fromJson(json)).toList();
      } else {
        // Log the error if the status code is not 200
        throw Exception('Failed to load data: ${response.statusCode}');
      }
    } catch (e) {
      // Log any exceptions that occur
      print('Error occurred: $e');
      throw Exception('Failed to fetch data: $e');
    }
  }
}
