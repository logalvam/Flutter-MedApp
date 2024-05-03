import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:medapp/apimethods/loginHistoryM.dart';
import 'package:medapp/colors.dart';
import '../apimethods/loginvalidationM.dart';
import '../colors.dart';
import '../dashboard.dart';
import '../method/validationM.dart';
import '../method/dashboardM.dart';
import 'package:shared_preferences/shared_preferences.dart';

class Loginpage extends StatefulWidget {
  const Loginpage({super.key});

  @override
  State<Loginpage> createState() => _LoginpageState();
}

String role = "";
Autentication? Result;
LoginHisInsertA logHis = LoginHisInsertA();

class _LoginpageState extends State<Loginpage> {
  TextEditingController _emailController = TextEditingController(text: "");
  TextEditingController _passwordController = TextEditingController(text: "");
  var auth = autenticationToJson;
  validationM loginV = validationM();
  @override
  void setState(VoidCallback fn) {
    // TODO: implement setState

    super.setState(fn);
  }

  @override
  void initState() {
    assigndetails();
    // TODO: implement initState
    super.initState();
  }

  assigndetails() async {
    SharedPreferences sref = await SharedPreferences.getInstance();
    // role = sref.getString("role")!;
    sref.setString("role", "");
    sref.setString("user", "");
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;

    return Scaffold(
      appBar: AppBar(
        title: Text("MEDAPP"),
        centerTitle: true,
        toolbarHeight: 70,
        backgroundColor: AppColors.navBarColor,
        titleTextStyle: TextStyle(
          color: Colors.white,
          fontSize: 40,
        ),
      ),
      body: Container(
        color: AppColors.bgColor,
        width: screenWidth,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Card(
              elevation: 15,
              child: Column(
                children: [
                  Text(
                    "Login",
                    style: TextStyle(fontSize: screenWidth * 0.09),
                  ),
                  Container(
                    decoration:
                        BoxDecoration(borderRadius: BorderRadius.circular(40)),
                    width: screenWidth * 0.45,
                    margin: EdgeInsets.fromLTRB(45, 10, 45, 0),
                    child: TextField(
                      controller: _emailController,
                      decoration: InputDecoration(
                        hintText: "Email",
                        hintStyle: TextStyle(
                            color: Color.fromARGB(255, 143, 117, 117)),
                        border: UnderlineInputBorder(),
                      ),
                    ),
                  ),
                  Container(
                    width: screenWidth * 0.45,
                    child: TextFormField(
                      controller: _passwordController,
                      obscureText: true,
                      decoration: InputDecoration(
                        hintText: "Password",
                        hintStyle: TextStyle(
                            color: Color.fromARGB(255, 143, 117, 117)),
                        border: UnderlineInputBorder(),
                      ),
                    ),
                  ),
                  MaterialButton(
                    onPressed: () async {
                      SharedPreferences sref =
                          await SharedPreferences.getInstance();

                      var email = _emailController.text;
                      var pass = _passwordController.text;
                      if (email.isEmpty || pass.isEmpty) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                            backgroundColor: Colors.red,
                            elevation: 5,
                            // animation: Animation,
                            content: Container(
                              child: Text("Please fill all the fields"),
                            ),
                          ),
                        );
                      } else if (!email.contains("@") ||
                          !email.contains(".com")) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                            backgroundColor: Colors.red,
                            elevation: 5,
                            // animation: Animation,
                            content: Container(
                              child: Text("Please enter the vlaid email"),
                            ),
                          ),
                        );
                      } else if (pass.length < 8) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                            backgroundColor: Colors.red,
                            elevation: 5,
                            // animation: Animation,
                            content: Container(
                              child: Text("not a valid Password"),
                            ),
                          ),
                        );
                      } else if (!email.isEmpty && !pass.isEmpty) {
                        Result = await loginV.loginvalidate(email, pass);
                        print(Result!.role);
                        var isInsert = await logHis.LoginHisInsertApi(email);
                        print("isInsert $isInsert");

                        if (Result!.role.isNotEmpty) {
                          sref.setString("role", Result!.role);
                          sref.setString("user", email);
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => DashBoard()));
                          print("user");
                        } else {
                          ScaffoldMessenger.of(context).showSnackBar(
                            SnackBar(
                              backgroundColor: Colors.red,
                              elevation: 5,
                              // animation: Animation,
                              content: Container(
                                child: Text("Invalid Username And Password"),
                              ),
                            ),
                          );
                        }
                      }
                    },
                    padding: EdgeInsets.all(20),
                    child:
                        // Result!.toJson().isNotEmpty
                        //     ? Center(
                        //         child: CircularProgressIndicator(),
                        //       )
                        Container(
                      width: screenWidth * 0.2,
                      child: Row(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Icon(Icons.login),
                          SizedBox(width: 10),
                          Text(
                            "Login",
                            style: TextStyle(
                                color: Color.fromARGB(255, 2, 2, 2),
                                fontSize: screenWidth * 0.03),
                          )
                        ],
                      ),
                    ),
                  ),
                  // MaterialButton(
                  //     onPressed: () {},
                  //     child: SizedBox(
                  //       child: RichText(
                  //           text: TextSpan(children: [
                  //         TextSpan(
                  //           text: "Don't have an account? ",
                  //           style: TextStyle(
                  //             color: Color.fromARGB(255, 143, 117, 117),
                  //             fontSize: 18,
                  //           ),
                  //         ),
                  //         TextSpan(
                  //           text: "Sign Up",
                  //           style: TextStyle(
                  //             color: Color.fromARGB(255, 143, 117, 117),
                  //             fontSize: 18,
                  //           ),
                  //         ),
                  //       ])),
                  //     )),
                  SizedBox(
                    height: screenWidth * 0.05,
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
