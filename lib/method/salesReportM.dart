// To parse this JSON data, do
//
//     final salesreport = salesreportFromJson(jsonString);

import 'dart:convert';

Salesreport salesreportFromJson(String str) =>
    Salesreport.fromJson(json.decode(str));

String salesreportToJson(Salesreport data) => json.encode(data.toJson());

class Salesreport {
  List<Salesarr> salesarr;

  Salesreport({
    required this.salesarr,
  });

  factory Salesreport.fromJson(Map<String, dynamic> json) => Salesreport(
        salesarr: List<Salesarr>.from(
            json["salesarr"].map((x) => Salesarr.fromJson(x))),
      );

  Map<String, dynamic> toJson() => {
        "salesarr": List<dynamic>.from(salesarr.map((x) => x.toJson())),
      };
}

class Salesarr {
  int billno;
  DateTime billdate;
  String medname;
  int quantity;
  int amount;

  Salesarr({
    required this.billno,
    required this.billdate,
    required this.medname,
    required this.quantity,
    required this.amount,
  });

  factory Salesarr.fromJson(Map<String, dynamic> json) => Salesarr(
        billno: json["billno"],
        billdate: DateTime.parse(json["billdate"]),
        medname: json["medname"],
        quantity: json["quantity"],
        amount: json["amount"],
      );

  Map<String, dynamic> toJson() => {
        "billno": billno,
        "billdate":
            "${billdate.year.toString().padLeft(4, '0')}-${billdate.month.toString().padLeft(2, '0')}-${billdate.day.toString().padLeft(2, '0')}",
        "medname": medname,
        "quantity": quantity,
        "amount": amount,
      };
}
