import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/singleMedicineBillM.dart';
import '../method/stockviewM.dart';
import '../method/validationM.dart';

class SinglemedA {
  SingleMedM? model1;
  Future<dynamic> singleMedicineApi(String medname, int quantity) async {
    print("object");
    try {
      print(medname);
      print(quantity.runtimeType);
      var response =
          await http.put(Uri.parse("http://192.168.2.139:8260/MedicineInput"),
              body: jsonEncode({
                "medname": medname,
                "quantity": quantity,
              }));
      print(response.body);
      if (response.statusCode == 200) {
        model1 = await singleMedMFromJson(response.body);
        return model1?.medlist[0];
      } else {
        throw Exception(
            'Failed to load data. Status code: ${response.statusCode}');
      }
    } catch (e) {
      print('Error: $e');
    }
  }
}
