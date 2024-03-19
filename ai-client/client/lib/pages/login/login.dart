import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import '../../Controller/user_controller.dart';

class LoginRegisterPage extends StatefulWidget {
  @override
  _LoginRegisterPageState createState() => _LoginRegisterPageState();
}

class _LoginRegisterPageState extends State<LoginRegisterPage> {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  bool _isLogin = true;

  @override
  void initState() {
    super.initState();
    UserController.checkLogin();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(_isLogin ? 'Login' : 'Register'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            _buildTextField(_usernameController, 'Username', Icons.person_outline),
            const SizedBox(height: 20.0),
            _buildTextField(_passwordController, 'Password', Icons.lock_outline, isPassword: true),
            const SizedBox(height: 20.0),
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                foregroundColor: Colors.white, backgroundColor: Colors.grey[800], // Text color
              ),
              onPressed: () {
                String username = _usernameController.text;
                String password = _passwordController.text;
                if (_isLogin) {
                  UserController.login(username, password);
                } else {
                  UserController.register(username, password);
                }
              },
              child: Text(_isLogin ? 'Login' : 'Register'),
            ),
            SizedBox(height: 20.0),
            TextButton(
              style: TextButton.styleFrom(
                foregroundColor: Colors.grey[500], // Grey text button
              ),
              onPressed: () {
                setState(() {
                  _isLogin = !_isLogin;
                });
              },
              child: Text(_isLogin ? 'Create an account' : 'Already have an account? Login'),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildTextField(TextEditingController controller, String label, IconData icon, {bool isPassword = false}) {
    return TextField(
      cursorColor: Colors.grey,
      obscureText: isPassword,
      decoration: InputDecoration(
        label: Text(label, style: const TextStyle(color: Colors.black)),
        prefixIcon: Icon(icon),
        enabledBorder: const OutlineInputBorder(
          borderSide: BorderSide(color: Colors.grey),
        ),
        disabledBorder: const OutlineInputBorder(
          borderSide: BorderSide(color: Colors.grey),
        ),
        focusedBorder: const OutlineInputBorder(
          borderSide: BorderSide(color: Colors.black),
        ),
        focusColor: Colors.red,
      ),
      style: const TextStyle(),
      onChanged: (value) {
        controller.text = value;
      },
    );
  }
}
