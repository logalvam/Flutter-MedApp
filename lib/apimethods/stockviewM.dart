import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/stockviewM.dart';
import '../method/validationM.dart';

class StockViewA {
  StockviewM? model;
  Future<dynamic> stockViewApi() async {
    try {
      var response = await http.get(
        Uri.parse("http://192.168.2.139:8260/viewMedicine"),
      );
      if (response.statusCode == 200) {
        model = await stockviewFromJson(response.body);
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
