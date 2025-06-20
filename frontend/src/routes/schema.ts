import { z } from "zod/v4";
 
export const formSchema = z.object({
 startNumber: z.number().min(0, {
    message: "Numret måste vara minst 0"
 }),
 endNumber: z.number().min(0),
 documentType: z.enum(["Produktdatablad", "Säkerhetsdatablad"]).default("Produktdatablad"),
 sheetName: z.enum(["Dokument"]).default("Dokument"),
 fileName: z.string().nonempty({
  error: "Välj en fil"
 }),
}).refine((data) => data.endNumber > data.startNumber, {
  message: "Slutnummer måste vara större än startnumret",
  path: ["endNumber"], // the field that should show the error
});
 
export type FormSchema = typeof formSchema;