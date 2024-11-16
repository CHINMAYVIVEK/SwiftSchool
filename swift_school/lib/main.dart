import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter and Go Backend',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const DataListScreen(),
    );
  }
}

class DataListScreen extends StatefulWidget {
  const DataListScreen({super.key});

  @override
  _DataListScreenState createState() => _DataListScreenState();
}

class _DataListScreenState extends State<DataListScreen> {
  late Future<List<Data>> futureData;

  @override
  void initState() {
    super.initState();
    futureData = fetchData(); // Fetch the data when the widget is initialized
  }

  // Function to fetch data from the Go backend
  Future<List<Data>> fetchData() async {
    final response =
        await http.get(Uri.parse('http://localhost:8080/api/data'));

    if (response.statusCode == 200) {
      // If the server returns a 200 OK response, parse the JSON data
      List<dynamic> dataJson = json.decode(response.body);
      return dataJson.map((json) => Data.fromJson(json)).toList();
    } else {
      // If the server does not return a 200 OK response, throw an error
      throw Exception('Failed to load data');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Flutter and Go Backend'),
      ),
      body: FutureBuilder<List<Data>>(
        future: futureData,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text('Error: ${snapshot.error}'));
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(child: Text('No data available.'));
          }

          // Display the data in a ListView
          List<Data> data = snapshot.data!;
          return ListView.builder(
            itemCount: data.length,
            itemBuilder: (context, index) {
              return ListTile(
                title: Text(data[index].value),
                subtitle: Text('ID: ${data[index].id}'),
              );
            },
          );
        },
      ),
    );
  }
}

// Model to parse JSON data
class Data {
  final int id;
  final String value;

  Data({required this.id, required this.value});

  // Factory method to create a Data object from a JSON object
  factory Data.fromJson(Map<String, dynamic> json) {
    return Data(
      id: json['id'],
      value: json['value'],
    );
  }
}
