import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { BuyinsuranceInterface } from "../models/IBuyinsurance";
import { format } from 'date-fns'
import React from 'react';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Buyinsurances() {
  const classes = useStyles();
  const [buyinsurances, setbuyinsurances] = useState<BuyinsuranceInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getBuyinsurances = async () => {
    fetch(`${apiUrl}/buyinsurances`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setbuyinsurances(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getBuyinsurances();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการซื้อประกัน
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/buyinsurance/create"
              variant="contained"
              color="primary"
            >
              เพิ่มราย
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="15%">
                  ลูกค้า
                </TableCell>
                <TableCell align="center" width="15%">
                  ประเภท
                </TableCell>
                <TableCell align="center" width="15%">
                  พนักงานดูแล
                </TableCell>
                <TableCell align="center" width="15%">
                  วันที่ทำ
                </TableCell>
                <TableCell align="center" width="15%">
                  สุขภาพเบื้องต้น
                </TableCell>
                <TableCell align="center" width="15%">
                  ยอมรับนโยบาย
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {buyinsurances.map((item: BuyinsuranceInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Customer.Name}</TableCell>
                  <TableCell align="center">{item.InsuranceConverage.Protection.Protection_name}</TableCell>
                  <TableCell align="center">{item.Employee.Name}</TableCell>
                  <TableCell align="center">{format((new Date(item.Adddate)), 'dd MMMM yyyy ')}</TableCell>
                  <TableCell align="center">{item.Healthinfrom}</TableCell>
                  <TableCell align="center">{item.Consent}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Buyinsurances;