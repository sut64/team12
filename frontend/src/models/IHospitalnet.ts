import { EmployeesInterface } from "./IEmployee";
import { ProvinceInterface } from "./IProvince";
import { StatusInterface } from "./IStatus";
import { GenreInterface } from "./IGenre";

export interface HospitalnetInterface {
    ID: number,
    Name: string,
    Contract: number,
    Address: string,
    Adddate: string,

    EmployeeID: number,
    Employee: EmployeesInterface,

    StatusID: number,
    Status: StatusInterface,

    ProvinceID: number,
    Province: ProvinceInterface,

    GenreID: number,
    Genre: GenreInterface
}