import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/salesReportM.dart';
import '../method/stockviewM.dart';
import '../method/validationM.dart';

class stockValueA {
  var totalValue;
  Future<dynamic> stockValueApi() async {
    try {
      var response = await http.put(
        Uri.parse("http://192.168.2.139:8260/stockprices"),
      );
      if (response.statusCode == 200) {
        var ReturnValue = await jsonDecode(response.body);
        print("stock: $ReturnValue");
        totalValue = ReturnValue["Total"];
        return;
      } else {
        throw Exception(
            'Failed to load data. Status code: ${response.statusCode}');
      }
    } catch (e) {
      print('Error: $e');
    }
  }
}
