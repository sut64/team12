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

import { PackageInsurInterface } from "../models/IPackageInsur";
import { ProtectionInterface } from "../models/IProtection";
import { EmployeesInterface } from "../models/IEmployee";
import { TotallistInterface } from "../models/ITotallist";
import { InsuranceConverageInterface } from "../models/IInsuranceConverage";

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

function InsuranceConverageCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());;
  const [packageinsurs, setPackageInsurs] = useState<PackageInsurInterface[]>([]);
  const [protections, setProtections] = useState<ProtectionInterface[]>([]);
  const [totallists, setTotallists] = useState<TotallistInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeesInterface>();
  const [insuranceconverage, setInsuranceConverage] = useState<Partial<InsuranceConverageInterface>>(
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
    const name = event.target.name as keyof typeof insuranceconverage;
    setInsuranceConverage({
      ...insuranceconverage,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getPackageInsur = async () => {
    fetch(`${apiUrl}/packageinsurs`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setPackageInsurs(res.data);
        } else {
          console.log("else");
        }
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

 const getTotallist = async () => {
    fetch(`${apiUrl}/totallists`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setTotallists(res.data);
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
        insuranceconverage.EmployeeID = res.data.ID
        if (res.data) {
          setEmployees(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getPackageInsur();
    getProtection();
    getTotallist();
    getEmployee();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      //InsuranceConverageID: convertType(insuranceconverage.ID),
      PackageInsurID: convertType(insuranceconverage.PackageInsurID),
      EmployeeID: convertType(insuranceconverage.EmployeeID),
      DateTime: selectedDate,
      ProtectionID: convertType(insuranceconverage.ProtectionID),
      TotallistID: convertType(insuranceconverage.TotallistID),
      Protectiontype:insuranceconverage.Protectiontype,
      Premium:convertType(insuranceconverage.Premium),
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

    fetch(`${apiUrl}/insuranceconverages`, requestOptionsPost)
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
              บันทึกรายการความคุ้มครอง
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>รูปแบบความคุ้มครอง</p>
              <Select
                native
                value={insuranceconverage.ID}
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
              <p>รายการเหมาจ่ายค่ารักษารวม</p>
              <Select
                native
                value={insuranceconverage.ID}
                onChange={handleChange}
                inputProps={{
                  name: "TotallistID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรายการเหมาจ่ายค่ารักษารวม
                </option>
                {totallists.map((item: TotallistInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Totallist_cost}
                  </option>
                ))}
		</Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>รายการแพ็คเกจ</p>
              <Select
                native
                value={insuranceconverage.ID}
                onChange={handleChange}
                inputProps={{
                  name: "PackageInsurID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรายการแพ็คเกจ
                </option>
                {packageinsurs.map((item: PackageInsurInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Package_name}
                  </option>
                ))}
             
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="DateTime"
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
              <p>ใส่ประเภทการรักษา</p>
                <TextField
                  name="Protectiontype"
                  value={insuranceconverage.Protectiontype}
                  onChange={handleChange}
                  label="กรุณาระบุประเภทการรักษา" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ใส่เบี้ยประกันเริ่มต้นต่อเดือน</p>
                <TextField
                  name="Premium"
                  value={insuranceconverage.Premium}
                  onChange={handleChange}
                  label="กรุณาระบุเบี้ยเริ่มต้น" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>พนักงาน</p>
              <Select
                native
                value={insuranceconverage.EmployeeID}
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
              component={RouterLink}
              to="/insuranceconverages"
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
              บันทึกความคุ้มครอง
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default InsuranceConverageCreate;