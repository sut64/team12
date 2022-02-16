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
import { InsuranceClaimInterface } from "../models/IInsuranceClaim";
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

function InsuranceClaims() {
  const classes = useStyles();
  const [insuranceclaims, setInsuranceClaims] = useState<InsuranceClaimInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getInsuranceClaims = async () => {
    fetch(`${apiUrl}/insuranceclaims`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setInsuranceClaims(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getInsuranceClaims();
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
              ข้อมูลการเคลม
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/insuranceclaims_create"
              variant="contained"
              color="primary"
            >
              เพิ่มรายการเคลม
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
                  สาเหตุ
                </TableCell>
                <TableCell align="center" width="15%">
                  ลูกค้า
                </TableCell>
                <TableCell align="center" width="15%">
                  พนักงาน
                </TableCell>
                <TableCell align="center" width="15%">
                  ข้อมูลผู้ป่วย
                </TableCell>
                <TableCell align="center" width="15%">
                  ค่าสินไหมทดแทน
                </TableCell>
                <TableCell align="center" width="15%">
                  วันที่การเคลม
                </TableCell>
            
              </TableRow>
            </TableHead>
            <TableBody>
              {insuranceclaims.map((item: InsuranceClaimInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Motive.Name}</TableCell>
                  <TableCell align="center">{item.Customer.Name}</TableCell>
                  <TableCell align="center">{item.Employee.Name}</TableCell>
                  <TableCell align="center">{item.Patient}</TableCell>
                  <TableCell align="center">{item.Compensation}</TableCell>
                  <TableCell align="center">{format((new Date(item.Insdate)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default InsuranceClaims;