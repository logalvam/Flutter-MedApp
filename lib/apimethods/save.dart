import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/stockviewM.dart';
import '../method/validationM.dart';

//BillMaster Api
class savePurchasedItemsA {
  savePurchasedItemsApi(List billDetails) async {
    try {
      // print(email);
      var response = await http.put(
          Uri.parse("http://192.168.2.139:8260/MedicineOutput"),
          body: jsonEncode(billDetails));
      if (response.statusCode == 200) {
        Map resp = await jsonDecode(response.body);
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

//Bill Details Api
class BillDePurchasedItemsA {
  BillDePurchasedItemsApi(List billdetails) async {
    try {
      // print(email);
      var response = await http.put(
          Uri.parse("http://192.168.2.139:8260/BillDetails"),
          body: jsonEncode(billdetails));
      if (response.statusCode == 200) {
        Map resp = await jsonDecode(response.body);
        print(resp);
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
