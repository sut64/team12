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

import { InvoicesInterface } from "../models/IInvoice";
import { CustomersInterface } from "../models/ICustomer";
import { EmployeesInterface } from "../models/IEmployee";
import { InvoicePaymentInterface } from "../models/IInvoicePayment";

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

function InvoicePaymentCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());;
  const [invoices, setInvoices] = useState<InvoicesInterface[]>([]);
  const [customers, setCustomers] = useState<CustomersInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeesInterface>();
  const [invoicePayment, setInvoicePayment] = useState<Partial<InvoicePaymentInterface>>(
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
    const name = event.target.name as keyof typeof invoicePayment;
    setInvoicePayment({
      ...invoicePayment,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getInvoices = async () => {
    fetch(`${apiUrl}/invoices`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setInvoices(res.data);
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
        invoicePayment.EmployeeID = res.data.ID
        if (res.data) {
          setEmployees(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getInvoices();
    getCustomers();
    getEmployee();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      InvoiceID: convertType(invoicePayment.InvoiceID),
      CustomerID: convertType(invoicePayment.CustomerID),
      EmployeeID: convertType(invoicePayment.EmployeeID),
      PaymentTime: selectedDate,
      InvoiceNumber: invoicePayment.InvoiceNumber,
      PaymentAmount: convertType(invoicePayment.PaymentAmount),
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

    fetch(`${apiUrl}/invoice_payments`, requestOptionsPost)
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
    <Container className={classes.container} maxWidth="md">
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
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกรายการชำระ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>รูปแบบชำระ</p>
              <Select
                native
                value={invoicePayment.InvoiceID}
                onChange={handleChange}
                inputProps={{
                  name: "InvoiceID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรูปแบบชำระ
                </option>
                {invoices.map((item: InvoicesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.PaymentType}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ลูกค้า</p>
              <Select
                native
                value={invoicePayment.CustomerID}
                onChange={handleChange}
                inputProps={{
                  name: "CustomerID",
                }}
              >
                <option aria-label="None" value="">
                กรุณาเลือกลูกค้า
                </option>
                {customers.map((item: CustomersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>พนักงาน</p>
              <Select
                native
                value={invoicePayment.EmployeeID}
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
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="PaymentTime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เลขใบแจ้ง</p>
                <TextField
                  name="InvoiceNumber"
                  value={invoicePayment.InvoiceNumber}
                  onChange={handleChange}
                  label="กรุณาระบุเลขใบแจ้ง" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>จำนวน</p>
                <TextField
                  name="PaymentAmount"
                  value={invoicePayment.PaymentAmount}
                  onChange={handleChange}
                  label="กรุณาระบุจำนวนเงิน" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/invoice_payments"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึกการชำระ
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default InvoicePaymentCreate;