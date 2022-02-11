import { ProtectionInterface } from "./IProtection";
import {PackageInsurInterface } from "./IPackageInsur";
import { EmployeesInterface } from "./IEmployee";
import { TotallistInterface } from "./ITotallist";

export interface  InsuranceConverageInterface {
    ID: number,
    PackageInsurID: number,
    PackageInsur: PackageInsurInterface,
    ProtectionID: number,
    Protection: ProtectionInterface,
    TotallistID: number,
    Totallist: TotallistInterface,
    EmployeeID: number,
    Employee: EmployeesInterface,
    Protectiontype: string,
	Premium:        number,
	Datetime:      Date,

}    