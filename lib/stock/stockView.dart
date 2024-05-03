import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:medapp/colors.dart';
import 'package:medapp/method/stockviewM.dart';

import '../apimethods/stockviewM.dart';
import '../method/stockviewM.dart';

class StockViewd extends StatefulWidget {
  const StockViewd({super.key});

  @override
  State<StockViewd> createState() => _StockViewdState();
}

StockViewA Aview = StockViewA();

class _StockViewdState extends State<StockViewd> {
  StockviewM? Sviewd;
  @override
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

  Widget build(BuildContext context) {
    if (Sviewd!.medlist.isEmpty) {
      return Center(
        child: CircularProgressIndicator(), // Show loading indicator
      );
    }
    double screenWidth = MediaQuery.of(context).size.width;
    TextStyle fonts = TextStyle(
      fontSize: screenWidth * 0.04,
      fontFamily: "Helvetica",
    );
    return Scaffold(
        body: Container(
      color: AppColors.bgColor,
      child: ListView(
        children: [
          Column(
            children: [
              Container(
                width: screenWidth,
                child: DataTable(
                    columnSpacing: 3,
                    columns: <DataColumn>[
                      DataColumn(
                          tooltip: "Medicine Name",
                          label: Text(
                            "Medicine Name",
                            style: TextStyle(
                                fontSize: screenWidth * 0.03,
                                fontStyle: FontStyle.italic,
                                fontWeight: FontWeight.bold),
                          )),
                      DataColumn(
                          label: Text(
                        "Brand",
                        style: TextStyle(
                            fontSize: screenWidth * 0.03,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Quantity",
                        style: TextStyle(
                            fontStyle: FontStyle.italic,
                            fontSize: screenWidth * 0.03,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Price",
                        style: TextStyle(
                            fontSize: screenWidth * 0.03,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                    ],
                    rows: List<DataRow>.generate(
                        Sviewd!.medlist.length,
                        (index) => DataRow(cells: <DataCell>[
                              DataCell(
                                Text(
                                  Sviewd!.medlist[index].medname,
                                  style: fonts,
                                ),
                              ),
                              DataCell(
                                Text(
                                  Sviewd!.medlist[index].brand,
                                  style: fonts,
                                ),
                              ),
                              DataCell(Text(
                                  (Sviewd!.medlist[index].quantity).toString(),
                                  style: fonts)),
                              DataCell(Text(
                                (Sviewd!.medlist[index].unitprice).toString(),
                                style: fonts,
                              )),
                            ]))),
              ),
            ],
          ),
        ],
      ),
    ));
  }
}
