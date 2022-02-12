import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import React from "react";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import TextField from '@material-ui/core/TextField';
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { EmployeesInterface } from "../models/IEmployee";
import { ProvinceInterface } from "../models/IProvince";
import { GenreInterface } from "../models/IGenre";
import { StatusInterface } from "../models/IStatus";
import { HospitalnetInterface } from "../models/IHospitalnet";
import Hospitalnet from "./Hospitalnet";

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

function HospitalnetCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [employee, setEmployee] = useState<EmployeesInterface>();
  const [errorMessage, setErrorMessage] = useState("");
  const [genre, setGenre] = useState<GenreInterface[]>([]);
  const [status, setStatus] = useState<StatusInterface[]>([]);
  const [province, setProvince] = useState<ProvinceInterface[]>([]);
  const [hospitalnets, setHospitalnet] = useState<Partial<HospitalnetInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

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
    const name = event.target.name as keyof typeof hospitalnets;
    setHospitalnet({
      ...hospitalnets,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getGenre = async () => {
    fetch(`${apiUrl}/genre`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setGenre(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getEmployee = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/employee/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        hospitalnets.EmployeeID = res.data.ID
        if (res.data) {
          setEmployee(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getStatus = async () => {
    fetch(`${apiUrl}/status`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setStatus(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getProvince = async () => {
    fetch(`${apiUrl}/province`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setProvince(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getEmployee();
    getProvince();
    getGenre();
    getStatus();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      ProvinceID: convertType(hospitalnets.ProvinceID),
      EmployeeID: convertType(hospitalnets.EmployeeID),
      GenreID: convertType(hospitalnets.GenreID),
      StatusID: convertType(hospitalnets.StatusID),
      Adddate: selectedDate,
      Contract: convertType(hospitalnets.Contract),
      Address: hospitalnets.Address,
      Name: hospitalnets.Name,

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

    fetch(`${apiUrl}/hospitalnets`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          setErrorMessage("");
          
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
              บันทึกข้อมูลโรงพยาบาล
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3}  >
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภท</p>
              <Select
                native
                value={hospitalnets.GenreID}
                onChange={handleChange}
                inputProps={{
                  name: "GenreID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภท
                </option>
                {genre.map((genre: GenreInterface) => (
                  <option value={genre.ID} key={genre.ID}>
                    {genre.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>status</p>
              <Select
                native
                value={hospitalnets.StatusID}
                onChange={handleChange}
                inputProps={{
                  name: "StatusID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกStatus
                </option>
                {status.map((status: StatusInterface) => (
                  <option value={status.ID} key={status.ID}>
                    {status.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>จังหวัด</p>
              <Select
                native
                value={hospitalnets.ProvinceID}
                onChange={handleChange}
                inputProps={{
                  name: "ProvinceID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกจังหวัด
                </option>
                {province.map((province: ProvinceInterface) => (
                  <option value={province.ID} key={province.ID}>
                    {province.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่อ</p>
                <TextField
                  name="Name"
                  value={hospitalnets.Name}
                  onChange={handleChange}
                  label="กรุณาระบุชื่อ" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ที่อยู่</p>
                <TextField
                  name="Address"
                  value={hospitalnets.Address}
                  onChange={handleChange}
                  label="กรุณาระบุที่อยู่" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ระยะเวลาสัญญา</p>
                <TextField
                  name="Contract"
                  value={hospitalnets.Contract}
                  onChange={handleChange}
                  label="กรุณาระบุสัญญา" 
                />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDatePicker
                  name="adddate"
                  autoOk
                  value={selectedDate}
                  onChange={handleDateChange}
                  minDate={new Date("2018-01-01")}
                  format="yyyy/MM/dd"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>พนักงาน</p>
              <Select
                native
                value={hospitalnets.EmployeeID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                 กรุณาเลือกพนักงาน
                </option>
                <option value={employee?.ID} key={employee?.ID}>
                  {employee?.Name}
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
export default HospitalnetCreate;