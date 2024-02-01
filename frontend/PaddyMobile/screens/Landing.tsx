import { FC } from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import pages from "../constants/screens";

type Props = {
  navigation: any;
};

const LandingScreen: FC<Props> = ({ navigation }) => {
  return (
    <View style={styles.container}>
      <Image
        source={require("../assets/illustrations/water.gif")}
        style={styles.image}
      />
      <View style={{ position: "absolute", top: "40%", gap: 10 }}>
        <Text style={{ ...styles.appName, fontSize: 70 }}>कृषक</Text>
        <Text style={styles.appName}>नेपलिले नेपाल लाई </Text>
      </View>
      <View style={{ paddingVertical: 30 }}>
        <Text
          style={styles.btn}
          onPress={() => navigation.navigate(pages.signup)}
        >
          Get Started
        </Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    position: "relative",
    backgroundColor: "white",
  },
  image: {
    width: "100%",
    height: 450,
    opacity: 0.5,
  },
  appName: {
    fontSize: 40,
    alignSelf: "center",
    fontWeight: "500",
  },
  btn: {
    fontSize: 20,
    borderRadius: 30,
    backgroundColor: "#5a5",
    paddingVertical: 10,
    paddingHorizontal: 30,
    color: "white",
  },
});

export default LandingScreen;
