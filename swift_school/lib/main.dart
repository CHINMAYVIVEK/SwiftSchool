import 'package:flutter/material.dart';
import 'package:swift_school/screens/splash_screen.dart';

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
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      // home: const DataListScreen(),
      home: const SplashScreen(), // Set SplashScreen as the initial screen
    );
  }
}
