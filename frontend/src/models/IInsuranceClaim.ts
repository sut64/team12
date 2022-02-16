import { MotivesInterface } from "./IMotive";
import { EmployeesInterface } from "./IEmployee";
import { CustomersInterface } from "./ICustomer";

export interface InsuranceClaimInterface {
    ID: number,
    MotiveID: number,
    Motive: MotivesInterface,

    EmployeeID: number,
    Employee: EmployeesInterface,
    
    Insdate: Date,
    Compensation: number,
    Patient: string,

    CustomerID: number,
    Customer: CustomersInterface,

}    