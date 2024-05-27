import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import '../models/user.dart';
import '../models/token.dart';

class AuthController with ChangeNotifier {
  String _baseUrl = 'http://10.0.2.2:5001';

  Token get token => _token;

  Future<void> signup(User user, BuildContext context) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/signup/'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(user.toJson()),
    );

    if (response.statusCode == 200) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Cadastro realizado com sucesso')),
      );
      Navigator.pushReplacementNamed(context, '/');
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Falha ao realizar o cadastro')),
      );
    }
  }

  Future<void> login(String email, String password, BuildContext context) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/token'),
      headers: <String, String>{
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: {
        'username': email,
        'password': password,
      },
    );

    if (response.statusCode == 200) {
      _token = Token.fromJson(jsonDecode(response.body));
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Login realizado com sucesso')),
      );
      Navigator.pushReplacementNamed(context, '/home');
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Falha no login')),
      );
    }
    notifyListeners();
  }

  Future<User> getCurrentUser() async {
    if (_token == null) return null;

    final response = await http.get(
      Uri.parse('$_baseUrl/users/me/'),
      headers: <String, String>{
        'Authorization': 'Bearer ${_token.accessToken}',
      },
    );

    if (response.statusCode == 200) {
      return User.fromJson(jsonDecode(response.body));
    } else {
      throw Exception('Falha ao carregar o usuário');
    }
  }

  void logout(BuildContext context) {
    _token = null;
    Navigator.pushReplacementNamed(context, '/');
    notifyListeners();
  }
}
