import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/salesReportM.dart';
import '../method/stockviewM.dart';
import '../method/validationM.dart';

class SalesReportA {
  Salesreport? model;
  Future<dynamic> SalesReportApi(String from, String to) async {
    try {
      var response = await http.put(
          Uri.parse("http://192.168.2.139:8260/salesreport"),
          body: jsonEncode({"date1": from, "date2": to}));
      if (response.statusCode == 200) {
        model = await salesreportFromJson(response.body);
        print("model: $model");
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
