import { InsuranceConverageInterface } from "./IInsuranceConverage"; 
import { CustomersInterface } from "./ICustomer"; 
import { EmployeesInterface } from "./IEmployee";

export interface BuyinsuranceInterface {
    ID: number,
    EmployeeID: number, 
    Employee: EmployeesInterface, 
    CustomerID: number, 
    Customer:CustomersInterface,
    Adddate: Date,
    InsuranceConverageID: number
    InsuranceConverage: InsuranceConverageInterface, 
    Consent: string,
    Healthinfrom: string,
}