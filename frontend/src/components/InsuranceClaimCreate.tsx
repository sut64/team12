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

import { MotivesInterface } from "../models/IMotive";
import { EmployeesInterface } from "../models/IEmployee";
import { CustomersInterface } from "../models/ICustomer";
import { InsuranceClaimInterface} from "../models/IInsuranceClaim";

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

function InsuranceClaimCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());;
  const [motives, setMotives] = useState<MotivesInterface[]>([]);
  const [customers, setCustomers] = useState<CustomersInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeesInterface>();
  const [insuranceclaime, setInsuranceClaim] = useState<Partial<InsuranceClaimInterface>>(
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
    const name = event.target.name as keyof typeof insuranceclaime;
    setInsuranceClaim({
      ...insuranceclaime,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getMotives = async () => {
    fetch(`${apiUrl}/motives`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setMotives(res.data);
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
        insuranceclaime.EmployeeID = res.data.ID
        if (res.data) {
          setEmployees(res.data);
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

  useEffect(() => {
    getMotives();
    getEmployee();
    getCustomers();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      MotiveID: convertType(insuranceclaime.MotiveID),
      CustomerID: convertType(insuranceclaime.CustomerID),
      EmployeeID: convertType(insuranceclaime.EmployeeID),
      Insdate: selectedDate,
      Compensation: convertType(insuranceclaime.Compensation),
      Patient: insuranceclaime.Patient ?? "",
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

    fetch(`${apiUrl}/insuranceclaim`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("???????????????????????????")
          setSuccess(true);
          setErrorMessage("")
        } else {
          console.log("????????????????????????????????????")
          setError(true);
          setErrorMessage(res.error)
        }
      });
  }

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          ??????????????????????????????????????????????????????
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          ???????????????????????????????????????????????????????????????: {errorMessage}
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
              ????????????????????????????????????????????????
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>??????????????????</p>
              <Select
                native
                value={insuranceclaime.MotiveID}
                onChange={handleChange}
                inputProps={{
                  name: "MotiveID",
                }}
              >
                <option aria-label="None" value="">
                  ????????????????????????????????????????????????
                </option>
                {motives.map((item: MotivesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>??????????????????</p>
              <Select
                native
                value={insuranceclaime.CustomerID}
                onChange={handleChange}
                inputProps={{
                  name: "CustomerID",
                }}
              >
                <option aria-label="None" value="">
                ????????????????????????????????????????????????
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
              <p>?????????????????????</p>
              <Select
                native
                value={insuranceclaime.EmployeeID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                  ???????????????????????????????????????????????????
                </option>
                <option value={employees?.ID} key={employees?.ID}>
                  {employees?.Name}
                </option>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>???????????????????????????????????????</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="Insdate"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="?????????????????????????????????????????????????????????????????????"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>???????????????????????????????????????</p>
                <TextField
                  name="Patient"
                  value={insuranceclaime.Patient}
                  onChange={handleChange}
                  label="???????????????????????????????????? 13 ????????????" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>??????????????????????????????????????????</p>
                <TextField
                  name="Compensation"
                  value={insuranceclaime.Compensation}
                  onChange={handleChange}
                  label="??????????????????????????????????????????????????????" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/insuranceclaims"
              variant="contained"
            >
              ????????????
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              ???????????????????????????????????????
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default InsuranceClaimCreate;