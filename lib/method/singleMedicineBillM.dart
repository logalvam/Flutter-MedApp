// To parse this JSON data, do
//
//     final singleMedM = singleMedMFromJson(jsonString);

import 'dart:convert';

SingleMedM singleMedMFromJson(String str) =>
    SingleMedM.fromJson(json.decode(str));

String singleMedMToJson(SingleMedM data) => json.encode(data.toJson());

class SingleMedM {
  List<Medlist> medlist;

  SingleMedM({
    required this.medlist,
  });

  factory SingleMedM.fromJson(Map<String, dynamic> json) => SingleMedM(
        medlist:
            List<Medlist>.from(json["medlist"].map((x) => Medlist.fromJson(x))),
      );

  Map<String, dynamic> toJson() => {
        "medlist": List<dynamic>.from(medlist.map((x) => x.toJson())),
      };
}

class Medlist {
  String medname;
  String brand;
  int quantity;
  int amount;

  Medlist({
    required this.medname,
    required this.brand,
    required this.quantity,
    required this.amount,
  });

  factory Medlist.fromJson(Map<String, dynamic> json) => Medlist(
        medname: json["medname"],
        brand: json["brand"],
        quantity: json["quantity"],
        amount: json["amount"],
      );

  Map<String, dynamic> toJson() => {
        "medname": medname,
        "brand": brand,
        "quantity": quantity,
        "amount": amount,
      };
}
