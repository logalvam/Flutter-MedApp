import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/loginhistoryM.dart';
import '../method/stockviewM.dart';
import '../method/validationM.dart';

class LoginHisViewA {
  LoginHisView? model;
  Future<dynamic> LoginHisViewApi() async {
    try {
      var response = await http.get(
        Uri.parse("http://192.168.2.139:8260/viewLogHistory"),
      );
      if (response.statusCode == 200) {
        model = await loginHisViewFromJson(response.body);
        return model?.hislist;
      } else {
        throw Exception(
            'Failed to load data. Status code: ${response.statusCode}');
      }
    } catch (e) {
      print('Error: $e');
    }
  }
}
