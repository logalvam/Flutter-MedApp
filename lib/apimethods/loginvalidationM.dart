import 'dart:convert';

import 'package:http/http.dart' as http;

import '../method/validationM.dart';

class validationM {
  Autentication? model;
  dynamic loginvalidate(String email, String pass) async {
    try {
      var response = await http.put(
          Uri.parse("http://192.168.2.139:8260/loginValidation"),
          body: json.encode({"userid": email, "password": pass}));
      if (response.statusCode == 200) {
        model = autenticationFromJson(response.body);

        return model;
      } else {
        throw Exception(
            'Failed to load data. Status code: ${response.statusCode}');
      }
    } catch (e) {
      print('Error: $e');
    }
  }
}
