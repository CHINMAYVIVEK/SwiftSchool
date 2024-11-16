// splash_screen.dart
import 'package:flutter/material.dart';
import 'package:swift_school/config/config.dart';
import 'package:swift_school/screens/data_list_screen.dart';
import 'dart:async';
// import 'home_page.dart'; // Import the home page

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  _SplashScreenState createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    // Simulate a delay for splash screen
    Timer(const Duration(seconds: 3), () {
      Navigator.pushReplacement(
        context,
        // MaterialPageRoute(builder: (context) => const HomePage()),
        MaterialPageRoute(builder: (context) => const DataListScreen()),
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.blueAccent,
      body: Center(
        child: AnimatedOpacity(
          opacity: 1.0,
          duration: const Duration(seconds: 2),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Container(
                padding: const EdgeInsets.all(40),
                decoration: BoxDecoration(
                  color: Colors.white.withOpacity(0.8),
                  borderRadius: BorderRadius.circular(20),
                  boxShadow: const [
                    BoxShadow(
                      color: Colors.black26,
                      offset: Offset(0, 4),
                      blurRadius: 10,
                    ),
                  ],
                ),
                child: const Column(
                  children: [
                    Icon(
                      Icons.attach_money,
                      size: 80,
                      color: Colors.green,
                    ),
                    SizedBox(height: 20),
                    Text(
                      AppConfig.appName,
                      style: TextStyle(
                        fontSize: 36,
                        fontWeight: FontWeight.bold,
                        color: Colors.blueAccent,
                        letterSpacing: 2,
                      ),
                    ),
                    SizedBox(height: 10),
                    CircularProgressIndicator(
                      valueColor:
                          AlwaysStoppedAnimation<Color>(Colors.blueAccent),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
