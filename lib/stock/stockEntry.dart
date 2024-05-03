import 'dart:collection';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter/widgets.dart';
import 'package:medapp/colors.dart';
import 'package:medapp/method/stockviewM.dart';
import 'package:medapp/stock/stockView.dart';

import '../apimethods/stockEntry.dart'; // Assuming this is the correct import

class StockEntry extends StatefulWidget {
  const StockEntry({Key? key}) : super(key: key);

  @override
  State<StockEntry> createState() => _StockEntryState();
}

StockviewM? Sviewd;

var Medicne_Master = {};

class _StockEntryState extends State<StockEntry> {
  String? selectedMedicine;
  String? selectedBrand;
  void initState() {
    Sviewd = StockviewM(medlist: []);
    fetchData();
    // TODO: implement initState
    super.initState();
  }

  Future<void> fetchData() async {
    try {
      await Aview.stockViewApi();
      setState(() {
        Sviewd = Aview.model;
      }); // Update the UI after fetching data
    } catch (e) {
      print('Error fetching data: $e');
    }
  }

  StockEntryA entry = StockEntryA();
  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    // double screenHeight = MediaQuery.of(context).size.height;
    TextEditingController? _brandLable = TextEditingController();
    TextEditingController _unitPrice = TextEditingController();
    TextEditingController _quantity = TextEditingController();
    TextEditingController? _medicineLale = TextEditingController();

    ScaffoldMsg(String message) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          backgroundColor: Color.fromARGB(255, 43, 41, 41),
          content: RichText(
              text: TextSpan(children: [
            TextSpan(
              text: message,
              style: TextStyle(
                  fontSize: screenWidth * 0.03,
                  color: Color.fromARGB(255, 255, 255, 255)),
            ),
          ])),
          duration: Duration(seconds: 3),
        ),
      );
    }

    AddBrandInStock() {
      if (_medicineLale!.text.contains(" ") ||
          _brandLable!.text.contains(" ") ||
          _brandLable!.text.isEmpty ||
          _medicineLale!.text.isEmpty) {
        ScaffoldMsg("Please Enter the details");
      } else {
        print("insert");
        _medicineLale.text = "";
        _brandLable.text = "";
        setState(() {});
      }
    }

    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Container(
              color: AppColors.bgColor,
              width: screenWidth * 0.6,
              child: Column(
                children: [
                  Container(
                    margin: EdgeInsets.fromLTRB(0, 25, 0, 0),
                    width: screenWidth * 0.5,
                    child: DropdownButtonFormField<String>(
                      value: selectedMedicine,
                      onChanged: (String? newValue) {
                        setState(() {
                          selectedMedicine = newValue!;
                          Map<String, String> Medicne_Master = {
                            for (var medicine in Sviewd!.medlist)
                              medicine.medname: medicine.brand
                          };

                          selectedBrand = Medicne_Master[selectedMedicine];
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
                    width: screenWidth * 0.5,
                    child: TextField(
                      enabled: false,
                      controller: _brandLable,
                      decoration: InputDecoration(
                          hintText: selectedBrand,
                          border: UnderlineInputBorder(
                            borderSide: BorderSide(
                              color: Colors.black,
                              width: 1.0,
                            ),
                          )),
                    ),
                  ),
                  Container(
                    width: screenWidth * 0.5,
                    child: TextField(
                      controller: _quantity,
                      inputFormatters: [
                        FilteringTextInputFormatter.deny(
                            RegExp(r'[a-zA-Z]')), // Deny numbers
                      ],
                      decoration: InputDecoration(
                          hintText: "Quantity",
                          border: UnderlineInputBorder(
                            borderSide: BorderSide(
                              color: Colors.black,
                              width: 1.0,
                            ),
                          )),
                    ),
                  ),
                  Container(
                    width: screenWidth * 0.5,
                    child: TextField(
                      controller: _unitPrice,
                      decoration: InputDecoration(
                          hintText: "Unit Price",
                          border: UnderlineInputBorder(
                            borderSide: BorderSide(
                              color: Colors.black,
                              width: 1.0,
                            ),
                          )),
                    ),
                  ),
                  Row(
                    children: [
                      MaterialButton(
                        onPressed: () {
                          try {
                            num? gQuantity = num.tryParse(_quantity.text);
                            num? gUnitPrice = num.tryParse(_unitPrice.text);
                            if (selectedMedicine == null) {
                              ScaffoldMsg("Please Select Medicine Name");
                            } else if (selectedBrand == null) {
                              ScaffoldMsg("Please fill all the fields");
                            } else if (_quantity.text.isEmpty ||
                                gQuantity!.isNegative ||
                                gQuantity.runtimeType == double) {
                              ScaffoldMsg("Please fill Corect Quantity");
                            } else if (_unitPrice.text.isEmpty ||
                                gUnitPrice!.isNegative) {
                              ScaffoldMsg(
                                  "Please insert correct Format UnitPrice");
                            } else {
                              var response = entry.StockEntryAPI(
                                  selectedMedicine!, gQuantity, gUnitPrice);
                              if (response != null) {
                                ScaffoldMsg("Stock Inerted");
                              }
                              // for (var i = 0; i < Stock.length; i++) {
                              //   if (Stock[i].medicine_name ==
                              //       selectedMedicine) {
                              //     Stock[i].unit_price = _unitPrice.text;
                              //     Stock[i].quantity =
                              //         (int.parse(Stock[i].unit_price) +
                              //                 int.parse(_quantity.text))
                              //             .toString();
                              //   }
                              // }
                              // print("inserted");

                              selectedBrand = null;
                              selectedMedicine = null;

                              _unitPrice.clear();
                              _quantity.clear();
                              setState(() {});
                            }
                          } catch (e) {
                            ScaffoldMsg("Don't insert Character");
                          }
                        },
                        child: Container(
                          margin: EdgeInsets.fromLTRB(0, 20, 0, 20),
                          padding: EdgeInsets.fromLTRB(10, 5, 10, 5),
                          decoration: BoxDecoration(
                            color: AppColors.buttonColor,
                            borderRadius: BorderRadius.circular(20),
                          ),
                          child: Text("Add"),
                        ),
                      ),
                      MaterialButton(
                        onPressed: () {
                          showDialog(
                            context: context,
                            builder: (BuildContext context) {
                              return SimpleDialog(
                                backgroundColor: AppColors.bgColor,
                                title: Text('Add  New Medicine'),
                                children: [
                                  Container(
                                      child: Column(children: [
                                    Container(
                                      width: screenWidth * 0.5,
                                      child: TextField(
                                        inputFormatters: [
                                          FilteringTextInputFormatter.deny(
                                              RegExp(r'\d')), // Deny numbers
                                        ],
                                        controller: _medicineLale,
                                        decoration: InputDecoration(
                                            hintText: "Medicine Name",
                                            border: UnderlineInputBorder(
                                              borderSide: BorderSide(
                                                color: Colors.black,
                                                width: 1.0,
                                              ),
                                            )),
                                      ),
                                    ),
                                    Container(
                                      width: screenWidth * 0.5,
                                      child: TextField(
                                        controller: _brandLable,
                                        inputFormatters: [
                                          FilteringTextInputFormatter.deny(
                                              RegExp(r'\d')), // Deny numbers
                                        ],
                                        decoration: InputDecoration(
                                            hintText: "Brand",
                                            border: UnderlineInputBorder(
                                              borderSide: BorderSide(
                                                color: Colors.black,
                                                width: 1.0,
                                              ),
                                            )),
                                      ),
                                    ),
                                    Container(
                                      margin: EdgeInsets.fromLTRB(0, 20, 0, 10),
                                      child: Row(
                                        mainAxisAlignment:
                                            MainAxisAlignment.center,
                                        crossAxisAlignment:
                                            CrossAxisAlignment.center,
                                        children: [
                                          TextButton(
                                            onPressed: () {
                                              AddBrandInStock();
                                            },
                                            child: Text(
                                              "Insert",
                                              style: TextStyle(
                                                  color: Colors.black),
                                            ),
                                          ),
                                          TextButton(
                                            onPressed: () {
                                              Navigator.pop(context);
                                            },
                                            child: Text(
                                              "close",
                                              style: TextStyle(
                                                  color: Colors.black),
                                            ),
                                          ),
                                        ],
                                      ),
                                    )
                                  ])),
                                ],
                              );
                            },
                          );
                        },
                        child: Container(
                          padding: EdgeInsets.fromLTRB(10, 5, 10, 5),
                          margin: EdgeInsets.fromLTRB(0, 20, 0, 20),
                          // padding: EdgeInsets.all(5),
                          decoration: BoxDecoration(
                            color: AppColors.buttonColor,
                            borderRadius: BorderRadius.circular(20),
                          ),
                          child: Text("Add New"),
                        ),
                      )
                    ],
                  )
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
