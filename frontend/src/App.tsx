import React, { useEffect } from "react";
import clsx from "clsx";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import List from "@material-ui/core/List";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import Button from "@material-ui/core/Button";

import MonetizationOnIcon from '@material-ui/icons/MonetizationOn';
import LocalHospitalIcon from '@material-ui/icons/LocalHospital';
import SecurityIcon from '@material-ui/icons/Security';
import AssignmentLateIcon from '@material-ui/icons/AssignmentLate';
import LocationOnIcon from '@material-ui/icons/AddLocation';
import PaymentIcon from '@material-ui/icons/Payment';


import InvoicePayments from "./components/InvoicePayment";
import InvoicePaymentCreate from "./components/InvoicePaymentCreate";
import InsuranceClaims from "./components/InsuranceClaim";
import InsuranceClaimCreate from "./components/InsuranceClaimCreate";
import InsuranceConverage from "./components/InsuranceConverage";
import InsuranceConverageCreate from "./components/InsuranceConverageCreate";
import Buyinsurance from "./components/Buyinsurance";
import BuyinsuranceCreate from "./components/BuyinsuranceCreate";
import hospitalnet from "./components/Hospitalnet";
import HospitalnetCreate from "./components/HospitalnetCreate";
import Payback from "./components/Payback";
import PaybackCreate from "./components/PaybackCreate";

import SignIn from "./components/SignIn";

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },
    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
    a: {
      textDecoration: "none",
      color: "inherit",
    },
  })
);

export default function MiniDrawer() {
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const [token, setToken] = React.useState<String>("");
  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const menu = [
    { name: "?????????????????????????????????????????????", icon: <MonetizationOnIcon />, path: "/invoice_payments" },
    { name: "?????????????????????????????????????????????????????????", icon: <AssignmentLateIcon />, path: "/buyinsurances" },
    { name: "?????????????????????????????????????????????", icon: <PaymentIcon />, path: "/payback" },
    { name: "??????????????????????????????", icon: <LocalHospitalIcon />, path: "/insuranceclaims" },
    { name: "????????????????????????????????????", icon: <SecurityIcon />, path: "/insuranceconverage" },
    { name: "??????????????????????????????????????????????????????????????????????????????", icon: <LocationOnIcon />, path: "/hospitalnet" },
    
  ];

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  return (
    <div className={classes.root}>
      <Router>
        <CssBaseline />
        {token && (
          <>
            <AppBar
              position="fixed"
              className={clsx(classes.appBar, {
                [classes.appBarShift]: open,
              })}
            >
              <Toolbar>
                <IconButton
                  color="inherit"
                  aria-label="open drawer"
                  onClick={handleDrawerOpen}
                  edge="start"
                  className={clsx(classes.menuButton, {
                    [classes.hide]: open,
                  })}
                >
                  <MenuIcon />
                </IconButton>
                <Typography variant="h6" className={classes.title}>
                  G12 - ??????????????????????????????
                </Typography>
                <Button color="inherit" onClick={signout}>
                  ??????????????????????????????
                </Button>
              </Toolbar>
            </AppBar>
            <Drawer
              variant="permanent"
              className={clsx(classes.drawer, {
                [classes.drawerOpen]: open,
                [classes.drawerClose]: !open,
              })}
              classes={{
                paper: clsx({
                  [classes.drawerOpen]: open,
                  [classes.drawerClose]: !open,
                }),
              }}
            >
              <div className={classes.toolbar}>
                <IconButton onClick={handleDrawerClose}>
                  {theme.direction === "rtl" ? (
                    <ChevronRightIcon />
                  ) : (
                    <ChevronLeftIcon />
                  )}
                </IconButton>
              </div>
              <Divider />
              <List>
                {menu.map((item, index) => (
                  <Link to={item.path} key={item.name} className={classes.a}>
                    <ListItem button>
                      <ListItemIcon>{item.icon}</ListItemIcon>
                      <ListItemText primary={item.name} />
                    </ListItem>
                  </Link>
                ))}
              </List>
            </Drawer>
          </>
        )}

        <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
            <Switch>
              <Route exact path="/invoice_payments" component={InvoicePayments} />
              <Route exact path="/hospitalnet" component={hospitalnet} />
              <Route
                exact
                path="/invoice_payment/create"
                component={InvoicePaymentCreate}
              />
              <Route exact path="/insuranceclaims" component={InsuranceClaims} />
              <Route exact path="/insuranceclaims_create" component={InsuranceClaimCreate} />
              <Route exact path="/insuranceconverage" component={InsuranceConverage} />
              <Route exact path="/buyinsurances" component={Buyinsurance} />
              <Route exact path="/payback" component={Payback}/>
              <Route
                exact
                path="/insuranceconverage/create"
                component={InsuranceConverageCreate}
              />
              <Route
                exact
                path="/hospitanet/create"
                component={HospitalnetCreate}
              />
              <Route
                exact
                path="/buyinsurance/create"
                component={BuyinsuranceCreate}
              />
              <Route
                exact
                path="/payback/create"
                component={PaybackCreate}
              />
            </Switch>
          </div>
        </main>
      </Router>
    </div>
  );
}