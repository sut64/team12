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
import { InsuranceConverageInterface } from "../models/IInsuranceConverage";
import { format } from 'date-fns'

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

function InsuranceConverage() {
  const classes = useStyles();
  const [insuranceconverage, setinsuranceconverage] = useState<InsuranceConverageInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getInsuranceConverage = async () => {
    fetch(`${apiUrl}/insuranceconverages`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setinsuranceconverage(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getInsuranceConverage();
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
              ข้อมูลความคุ้มครอง
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/insuranceconverage/create"
              variant="contained"
              color="primary"
            >
              เพิ่มข้อมูลความคุ้มครอง
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  พนักงาน
                </TableCell>
                <TableCell align="center" width="15%">
                  ความคุ้มครอง
                </TableCell>
                <TableCell align="center" width="15%">
                  แพ็คเกจ
                </TableCell>
                <TableCell align="center" width="15%">
                  ทุนประกันภัยต่อปี
                </TableCell>
                <TableCell align="center" width="15%">
                  ทุนประกันภัยต่อเดือน
                </TableCell>
                <TableCell align="center" width="15%">
                  ประเภทของผู้ป่วย
                </TableCell>
                <TableCell align="center" width="15%">
                  วันที่ออกแพ็คเกจ
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {insuranceconverage.map((item: InsuranceConverageInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.Employee.Name}</TableCell>
                  <TableCell align="center">{item.Protection.Protection_name}</TableCell>
                  <TableCell align="center">{item.PackageInsur.Package_name}</TableCell>
                  <TableCell align="center">{item.Totallist.Totallist_cost}</TableCell> 
                  <TableCell align="center">{item.Premium}</TableCell>
                  <TableCell align="center">{item.Protectiontype}</TableCell>
                  <TableCell align="center">{format((new Date(item.Datetime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                  
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default InsuranceConverage;