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
import { format } from 'date-fns';
import { PaybackInterface } from "../models/IPat";
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

function Payback() {
  const classes = useStyles();
  const [payback, setPayback] = useState<PaybackInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getPayback = async () => {
    fetch(`${apiUrl}/paybacks`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setPayback(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getPayback();
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
              ข้อมูลการคืนทุนประกัน
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/payback/create"
              variant="contained"
              color="primary"
            >
              เพิ่มรายการคืนทุนประกัน
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="8%">
                  ลูกค้า
                </TableCell>
                <TableCell align="center" width="15%">
                  รูปแบบความคุ้มครอง
                </TableCell>
                <TableCell align="center" width="15%">
                  ธนาคาร
                </TableCell>
                <TableCell align="center" width="15%">
                  เลขบัญชี
                </TableCell>
                <TableCell align="center" width="15%">
                  รหัสบัตรประชาชน
                </TableCell>
                <TableCell align="center" width="15%">
                  จำนวนปี
                </TableCell>
                <TableCell align="center" width="20%">
                  พนักงาน
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {payback.map((item: PaybackInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.CustomerID}</TableCell>
                  <TableCell align="center">{item.ProtectionID}</TableCell>
                  <TableCell align="center">{item.BankID}</TableCell>
                  <TableCell align="center">{item.Accout}</TableCell>
                  <TableCell align="center">{item.IDcard}</TableCell>
                  <TableCell align="center">{item.Year}</TableCell>
                  <TableCell align="center">{item.Employee.Name}</TableCell>
      
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Payback;