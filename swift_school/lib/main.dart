import 'package:flutter/material.dart';
import 'screens/data_list_screen.dart';

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
