import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import TextField from '@material-ui/core/TextField';
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { EmployeesInterface } from "../models/IEmployee";
import { CustomersInterface } from "../models/ICustomer";
import { BankInterface } from "../models/Bank";
import { PaybackInterface } from "../models/IPat";
import { ProtectionInterface } from"../models/IProtection";


import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function PaybackCreate() {
  const classes = useStyles();
  const [employees, setEmployees] = useState<EmployeesInterface>();
  const [customer, setCustomers] = useState<CustomersInterface[]>([]);
  const [protections, setProtections] = useState<ProtectionInterface[]>([]);
  const [banks, setBanks] = useState<BankInterface[]>([]);
  const [payback, setPayback] = useState<Partial<PaybackInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof payback;
    setPayback({
      ...payback,
      [name]: event.target.value,
    });
  };

  const getProtection = async () => {
    fetch(`${apiUrl}/protections`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setProtections(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getCustomers = async () => {
    fetch(`${apiUrl}/customers`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCustomers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getEmployee = async () => {
    let eid = localStorage.getItem("uid");
    fetch(`${apiUrl}/employee/${eid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        payback.EmployeeID = res.data.ID
        if (res.data) {
          setEmployees(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getBank = async () => {
    fetch(`${apiUrl}/banks`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setBanks(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getCustomers();
    getEmployee();
    getBank();
    getProtection();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      CustomerID: convertType(payback.CustomerID),  
      EmployeeID: convertType(payback.EmployeeID),
      BankID: convertType(payback.BankID),
      ProtectionID: convertType(payback.ProtectionID),

      IDcard: payback.IDcard,
      Year: convertType(payback.Year),
      Accout: payback.Accout,
    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/payback`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
          setErrorMessage(res.error)
        }
      });
  }

  return (
    <Container>
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
        </Alert>
      </Snackbar>
      <Paper>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกข้อมูลการคืนทุนประกัน
            </Typography>
          </Box>
        </Box>
        <Divider />
   
        <Grid container spacing={3}  >

            <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ลูกค้า</p>
              <Select
                native
                value={payback.CustomerID}
                onChange={handleChange}
                inputProps={{
                  name: "CustomerID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกสมาชิก
                </option>
                {customer.map((item: CustomersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>รูปแบบความคุ้มครอง</p>
              <Select
                native
                value={payback.ProtectionID}
                onChange={handleChange}
                inputProps={{
                  name: "ProtectionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรายการความคุ้มครอง
                </option>
                {protections.map((item: ProtectionInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Protection_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ธนาคาร</p>
              <Select
                native
                value={payback.BankID}
                onChange={handleChange}
                inputProps={{
                  name: "BankID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกธนาคาร
                </option>
                {banks.map((item: BankInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Bank_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid> 

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เลขบัญชี</p>
                <TextField
                  name="Accout"
                  value={payback.Accout}
                  onChange={handleChange}
                  label="กรอกเลขบัญชี" 
                />
            </FormControl>
          </Grid>         

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>รหัสบัตรประชาชน</p>
                <TextField
                  name="IDcard"
                  value={payback.IDcard}
                  onChange={handleChange}
                  label="รหัสบัตรประชาชน 13 หลัก" 
                />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>จำนวนปี</p>
                <TextField
                  name="Year"
                  value={payback.Year}
                  onChange={handleChange}
                  label="ระบุจำนวนปีที่คืนทุนประกัน" 
                />
            </FormControl>
          </Grid>         
                    
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>พนักงาน</p>
              <Select
                native
                value={payback.EmployeeID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                 กรุณาเลือกพนักงาน
                </option>
                <option value={employees?.ID} key={employees?.ID}>
                  {employees?.Name}
                </option>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
    );
}
export default PaybackCreate;