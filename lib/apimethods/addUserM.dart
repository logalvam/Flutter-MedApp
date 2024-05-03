import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:medapp/stock/stockView.dart';

import '../method/singleMedicineBillM.dart';
import '../method/stockviewM.dart';
import '../method/validationM.dart';

class AdduserA {
  // SingleMedM? model1;
  Future<dynamic> AdduserAPI(
      String userid, String password, String role) async {
    try {
      var response =
          await http.put(Uri.parse("http://192.168.2.139:8260/addNewUser"),
              body: jsonEncode({
                "userid": userid,
                "password": password,
                "role": role,
              }));
      print(response.body);

      if (response.statusCode == 200) {
        Map resp = jsonDecode(response.body);
        // model1 = await singleMedMFromJson(response.body);
        print("respo: $resp");
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
