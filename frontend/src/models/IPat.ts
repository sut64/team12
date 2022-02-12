import { BankInterface } from "./Bank";
import { CustomersInterface } from "./ICustomer";
import { EmployeesInterface } from "./IEmployee";
import { ProtectionInterface } from "./IProtection";


export interface PaybackInterface {
    ID: number;
    IDcard: string;
    Accout: string;
    Year : number;
    BankID: number;
    Bank: BankInterface;
    CustomerID: number,
    Customer: CustomersInterface,
    EmployeeID: number,
    Employee: EmployeesInterface,
    ProtectionID: number,
    Protection: ProtectionInterface,
    
}