import 'dart:convert';
import 'dart:math';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter/widgets.dart';
import 'package:medapp/colors.dart';
import 'package:medapp/stock/stockEntry.dart';
import 'package:medapp/stock/stockView.dart';
import 'package:getwidget/getwidget.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../apimethods/save.dart';
import '../apimethods/singleMedBillM.dart';
import '../method/stockviewM.dart';

class Billentry extends StatefulWidget {
  const Billentry({super.key});

  @override
  State<Billentry> createState() => _BillentryState();
}

List<Map<String, dynamic>> sortedList = [];
StockviewM? Sviewd;
Map TemBill = {};
List TemBillList = [];
SinglemedA bill = SinglemedA();
String? selectedMedicine;
DateTime now = new DateTime.now();
savePurchasedItemsA saveitem = savePurchasedItemsA();
BillDePurchasedItemsA billD = BillDePurchasedItemsA();
Random random = Random();

class _BillentryState extends State<Billentry> {
  late num total = 0;
  late num gst = 0;
  late num netpay = 0;
  var userid;
  void initState() {
    Sviewd = StockviewM(medlist: []);
    fetchData();
    // TODO: implement initState
    super.initState();
  }

  Future<void> fetchData() async {
    SharedPreferences sref = await SharedPreferences.getInstance();
    userid = sref.getString("user");
    try {
      await Aview.stockViewApi();
      setState(() {
        Sviewd = Aview.model;
      }); // Update the UI after fetching data
    } catch (e) {
      print('Error fetching data: $e');
    }
  }

  NewBillGenerated() {
    TemBillList.add(TemBill);
    // print(TemBillList);

    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    TextEditingController _quantity = TextEditingController();
    var billno = random.nextInt(10000).toString().padLeft(4, '0');
    var date = new DateTime(now.year, now.month, now.day)
        .toString()
        .replaceAll("00:00:00.000", "");

    return Scaffold(
      body: ListView(
        children: [
          Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Card(
                margin: EdgeInsets.fromLTRB(0, 30, 0, 0),
                elevation: 5,
                child: Container(
                  color: AppColors.bgColor,
                  width: screenWidth * 0.8,
                  child: Column(
                    children: [
                      Container(
                          margin: EdgeInsets.fromLTRB(0, 30, 0, 0),
                          child: Text(
                            "Bill",
                            style: TextStyle(fontSize: screenWidth * 0.04),
                          )),
                      Container(
                        margin: EdgeInsets.fromLTRB(15, 10, 15, 0),
                        child: DropdownButtonFormField<String>(
                          value: selectedMedicine,
                          onChanged: (String? newValue) {
                            setState(() {
                              selectedMedicine = newValue!;
                              Map<String, String> Medicne_Master = {
                                for (var medicine in Sviewd!.medlist)
                                  medicine.medname: medicine.brand
                              };
                            });
                          },
                          items: Sviewd!.medlist.map((medicine) {
                            return DropdownMenuItem<String>(
                              value: medicine.medname,
                              child: Text(medicine.medname),
                            );
                          }).toList(),
                        ),
                      ),
                      Container(
                        decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(40)),
                        // width: screenWidth * 0.7,
                        margin: EdgeInsets.fromLTRB(15, 10, 15, 0),
                        child: TextField(
                          controller: _quantity,
                          inputFormatters: [
                            FilteringTextInputFormatter.deny(
                                RegExp(r'[a-zA-Z]')),
                          ],
                          decoration: InputDecoration(
                            hintText: "Unit Price",
                            hintStyle: TextStyle(
                                color: Color.fromARGB(255, 143, 117, 117)),
                            border: UnderlineInputBorder(),
                          ),
                        ),
                      ),
                      Container(
                        margin: EdgeInsets.fromLTRB(0, 10, 0, 10),
                        child: TextButton(
                            style: ButtonStyle(
                              foregroundColor:
                                  MaterialStateProperty.all<Color>(Colors.blue),
                              overlayColor:
                                  MaterialStateProperty.resolveWith<Color?>(
                                (Set<MaterialState> states) {
                                  if (states.contains(MaterialState.hovered))
                                    return Colors.blue.withOpacity(0.04);
                                  if (states.contains(MaterialState.focused) ||
                                      states.contains(MaterialState.pressed))
                                    return Colors.blue.withOpacity(0.12);
                                  return null; // Defer to the widget's default.
                                },
                              ),
                            ),
                            onPressed: () async {
                              // var temptotal = 0;
                              num tempgst = 0;
                              num tempnetpay = 0;
                              var qty = _quantity.text;
                              dynamic AddedBill = await bill.singleMedicineApi(
                                  selectedMedicine!, int.parse(qty));

                              TemBill = {
                                "medname": bill.model1!.medlist[0].medname,
                                "quantity": bill.model1!.medlist[0].quantity,
                                "brand": bill.model1!.medlist[0].brand,
                                "amount": bill.model1!.medlist[0].amount,
                              };
                              NewBillGenerated();
                              Map<String, Map<String, dynamic>> mergedData = {};

                              for (var item in TemBillList) {
                                var medname = item["medname"];
                                if (mergedData.containsKey(medname)) {
                                  mergedData[medname]!["quantity"] +=
                                      item["quantity"];
                                  mergedData[medname]!["amount"] +=
                                      item["amount"];
                                  // total += bill.model1!.medlist[0].amount;
                                } else {
                                  mergedData[medname] = Map.from(item);
                                  // total += item["amount"];
                                }
                              }
                              sortedList = mergedData.values.toList();

                              setState(() {
                                total = sortedList.fold(
                                    0.0, (sum, med) => sum + med["amount"]);

                                gst = total * 0.18;

                                netpay = total + gst;
                              });
                              // NewBillGenerated();
                            },
                            child: Text(
                              'Add',
                              style: TextStyle(fontSize: screenWidth * 0.03),
                            )),
                      )
                    ],
                  ),
                ),
              ),
              Divider(
                color: Colors.grey,
              ),
              Row(
                children: [
                  MaterialButton(
                    onPressed: () {
                      showDialog(
                          context: context,
                          builder: (BuildContext context) {
                            return AlertDialog(
                              title: Row(
                                children: [
                                  Text("Bill Preview"),
                                  Spacer(
                                    flex: 1,
                                  ),
                                  Text("Total ${total}"),
                                  Spacer(
                                    flex: 1,
                                  ),
                                  Text("GST ${gst}"),
                                  Spacer(
                                    flex: 1,
                                  ),
                                  Text("Net Price ${netpay}"),
                                  Spacer(
                                    flex: 1,
                                  ),
                                  IconButton(
                                    onPressed: () {
                                      Navigator.of(context).pop();
                                    },
                                    icon: Icon(Icons.close),
                                  )
                                ],
                              ),
                              actions: [],
                              content: Container(
                                width: screenWidth * 0.8,
                                // color: AppColors.bgColor,
                                child: DataTable(
                                    columnSpacing: 4,
                                    columns: <DataColumn>[
                                      DataColumn(
                                          tooltip: "Medicine Name",
                                          label: Text(
                                            "Medicine Name",
                                            style: TextStyle(
                                                fontSize: screenWidth * 0.02,
                                                fontStyle: FontStyle.italic,
                                                fontWeight: FontWeight.bold),
                                          )),
                                      DataColumn(
                                          label: Text(
                                        "Quantity",
                                        style: TextStyle(
                                            fontStyle: FontStyle.italic,
                                            fontSize: screenWidth * 0.02,
                                            fontWeight: FontWeight.bold),
                                      )),
                                      DataColumn(
                                          label: Text(
                                        "Price",
                                        style: TextStyle(
                                            fontSize: screenWidth * 0.02,
                                            fontStyle: FontStyle.italic,
                                            fontWeight: FontWeight.bold),
                                      )),
                                    ],
                                    rows: List<DataRow>.generate(
                                        sortedList.length,
                                        (index) => DataRow(cells: <DataCell>[
                                              DataCell(
                                                Text(
                                                  sortedList[index]["medname"],
                                                ),
                                              ),
                                              DataCell(Text(
                                                sortedList[index]["quantity"]
                                                    .toString(),
                                              )),
                                              DataCell(Text(
                                                sortedList[index]["amount"]
                                                    .toString(),
                                              )),
                                            ]))),
                              ),
                            );
                          });
                    },
                    child: Container(
                      padding: EdgeInsets.all(10),
                      decoration: BoxDecoration(
                          border: Border.all(color: AppColors.buttonColor),
                          color: AppColors.navBarColor,
                          borderRadius: BorderRadius.circular(20)),
                      child: Center(
                        child: Text("Preview"),
                      ),
                    ),
                  ),
                  MaterialButton(
                    onPressed: () async {
                      // List Billmaster = [];
                      // for (var i = 0; i < sortedList.length; i++) {
                      //   var bgst = (sortedList[i]["amount"]) * 0.18;
                      //   var bnetpay = (sortedList[i]["amount"]) + bgst;
                      //   Billmaster.add({
                      //     "billno": billno,
                      //     "amount": total,
                      //     "billgst": bgst,
                      //     "netprice": bnetpay,
                      //     "userid": userid
                      //   });
                      // }
                      // print("billdetails: $Billmaster");
                      // print("SortLit = $sortedList");
                      // var BillMaster =
                      //     await saveitem.savePurchasedItemsApi(Billmaster);

                      List billDetails = [];

                      for (var i = 0; i < sortedList.length; i++) {
                        var amount = sortedList[i]["amount"];
                        print("Amount: $amount");
                        var qty = sortedList[i]["quantity"];
                        print("Quantity: $qty");

                        if (qty != 0) {
                          var unitprice = amount / qty;
                          print("Unit Price: $unitprice");

                          billDetails.add({
                            "billno": billno,
                            "medname": sortedList[i]["medname"],
                            "quantity": qty,
                            "unitprice": unitprice,
                            "amount": amount,
                            "userid": userid
                          });
                        } else {
                          print(
                              "Warning: Quantity is zero, skipping division.");
                        }
                      }
                      print("billDetails $billDetails");
                      var saved =
                          await billD.BillDePurchasedItemsApi(billDetails);
                      setState(() {});
                    },
                    child: Container(
                      padding: EdgeInsets.all(10),
                      decoration: BoxDecoration(
                          border: Border.all(color: AppColors.buttonColor),
                          color: AppColors.navBarColor,
                          borderRadius: BorderRadius.circular(20)),
                      child: Center(
                        child: Text("Save"),
                      ),
                    ),
                  )
                ],
              ),
              Row(
                children: [
                  Expanded(
                    child: Text("Billno: $billno", textAlign: TextAlign.center),
                  ),
                  Expanded(
                    child: Text("Date: $date", textAlign: TextAlign.center),
                  ),
                  Expanded(
                    child: Text("Total: ${total.toStringAsFixed(2)}",
                        textAlign: TextAlign.center),
                  ),
                  Expanded(
                    child: Text("Gst: ${gst.toStringAsFixed(2)}",
                        textAlign: TextAlign.center),
                  ),
                  Expanded(
                    child: Text("Net Pay: ${netpay.toStringAsFixed(2)}",
                        textAlign: TextAlign.center),
                  ),
                ],
              ),
              Container(
                width: screenWidth * 0.8,
                color: AppColors.bgColor,
                child: DataTable(
                    columnSpacing: 4,
                    columns: <DataColumn>[
                      DataColumn(
                          tooltip: "Medicine Name",
                          label: Text(
                            "Medicine Name",
                            style: TextStyle(
                                fontSize: screenWidth * 0.02,
                                fontStyle: FontStyle.italic,
                                fontWeight: FontWeight.bold),
                          )),
                      DataColumn(
                          label: Text(
                        "Brand",
                        style: TextStyle(
                            fontSize: screenWidth * 0.02,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Quantity",
                        style: TextStyle(
                            fontStyle: FontStyle.italic,
                            fontSize: screenWidth * 0.02,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Price",
                        style: TextStyle(
                            fontSize: screenWidth * 0.02,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                    ],
                    rows: List<DataRow>.generate(
                        sortedList.length,
                        (index) => DataRow(cells: <DataCell>[
                              DataCell(
                                Text(
                                  sortedList[index]["medname"],
                                ),
                              ),
                              DataCell(
                                Text(
                                  sortedList[index]["brand"],
                                ),
                              ),
                              DataCell(Text(
                                sortedList[index]["quantity"].toString(),
                              )),
                              DataCell(Text(
                                sortedList[index]["amount"].toString(),
                              )),
                            ]))),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
