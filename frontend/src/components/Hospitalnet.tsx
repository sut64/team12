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
import { HospitalnetInterface } from "../models/IHospitalnet";
import { format } from 'date-fns';

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

function Hospitalnet() {
  const classes = useStyles();
  const [hospitalnet, setHospitalnet] = useState<HospitalnetInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getHospitalnet = async () => {
    fetch(`${apiUrl}/hospitalnets`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setHospitalnet(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getHospitalnet();
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
              ข้อมูลการบันทึก
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/hospitanet/create"
              variant="contained"
              color="primary"
            >
              เพิ่มรายการ
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
                  ประเภท
                </TableCell>
                <TableCell align="center" width="15%">
                  สถานะ
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่อ
                </TableCell>
                <TableCell align="center" width="15%">
                  จังหวัด
                </TableCell>
                <TableCell align="center" width="15%">
                  สัญญา
                </TableCell>
                <TableCell align="center" width="15%">
                  พนักงานที่ดูแล
                </TableCell>
                <TableCell align="center" width="20%">
                  วันที่บันทึก
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {hospitalnet.map((item: HospitalnetInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Genre.Name}</TableCell>
                  <TableCell align="center">{item.Status.Name}</TableCell>
                  <TableCell align="center">{item.Name}</TableCell>
                  <TableCell align="center">{item.Province.Name}</TableCell>
                  <TableCell align="center">{item.Contract}</TableCell>
                  <TableCell align="center">{item.Employee.Name}</TableCell>
                  <TableCell align="center">{format((new Date(item.Adddate)), 'dd-MM-yyyy')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Hospitalnet;