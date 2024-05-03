// To parse this JSON data, do
//
//     final stockview = stockviewFromJson(jsonString);

import 'dart:convert';

StockviewM stockviewFromJson(String str) =>
    StockviewM.fromJson(json.decode(str));

String stockviewToJson(StockviewM data) => json.encode(data.toJson());

class StockviewM {
  List<Medlist> medlist;

  StockviewM({
    required this.medlist,
  });

  factory StockviewM.fromJson(Map<String, dynamic> json) => StockviewM(
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
  int unitprice;

  Medlist({
    required this.medname,
    required this.brand,
    required this.quantity,
    required this.unitprice,
  });

  factory Medlist.fromJson(Map<String, dynamic> json) => Medlist(
        medname: json["medname"],
        brand: json["brand"],
        quantity: json["quantity"],
        unitprice: json["unitprice"],
      );

  Map<String, dynamic> toJson() => {
        "medname": medname,
        "brand": brand,
        "quantity": quantity,
        "unitprice": unitprice,
      };
}
