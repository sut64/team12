import { InvoicesInterface } from "./IInvoice";
import { CustomersInterface } from "./ICustomer";
import { EmployeesInterface } from "./IEmployee";

export interface InvoicePaymentInterface {
    ID: number,
    InvoiceID: number,
    Invoice: InvoicesInterface,
    CustomerID: number,
    Customer: CustomersInterface,
    EmployeeID: number,
    Employee: EmployeesInterface,
    PaymentTime: Date,
    InvoiceNumber: string,
    PaymentAmount: number,
}    