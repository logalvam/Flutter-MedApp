import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/stockviewM.dart';
import '../method/validationM.dart';

class TodaySalesA {
  TodaySalesApi(String email) async {
    try {
      print(email);
      var response = await http.put(
          Uri.parse("http://192.168.2.139:8260/todaysales"),
          headers: {"USER": email});
      if (response.statusCode == 200) {
        Map resp = await jsonDecode(response.body);
        print("RESP: $resp");
        return resp;
      } else {
        throw Exception(
            'Failed to load data. Status code: ${response.statusCode}');
      }
    } catch (e) {
      print('Error: $e');
    }
  }
}
