// import { z } from 'zod/v4';

// export const formSchema = z
// 	.object({
// 		startNumber: z.number().min(0, {
// 			message: 'Numret måste vara minst 0'
// 		}),
// 		endNumber: z.number().min(0),
// 		documentType: z
// 			.enum(['Produktdatablad', 'Säkerhetsdatablad', 'Prestandadeklaration'])
// 			.default('Produktdatablad'),
// 		sheetName: z.enum(['Dokument', 'illbruck']).default('Dokument'),
// 		fileName: z.string().nonempty({
// 			error: 'Välj en fil'
// 		})
// 	})
// 	.refine((data) => data.endNumber > data.startNumber, {
// 		message: 'Slutnummer måste vara större än startnumret',
// 		path: ['endNumber'] // the field that should show the error
// 	});

// export type FormSchema = typeof formSchema;

import { z } from 'zod/v4';

export const formSchema = z
	.object({
		startNumber: z.number().min(0, {
			message: 'Numret måste vara minst 0'
		}),
		endNumber: z.number().min(0),
		documentType: z
			.enum(['Produktdatablad', 'Säkerhetsdatablad', 'Prestandadeklaration'])
			.default('Produktdatablad'),
		sheetName: z.string().default('Dokument'),
		// sheetName: z.enum(['Dokument', 'illbruck']).default('Dokument'),
		fileName: z.string().min(1, {
			message: 'Välj en fil'
		}),
		articleNumberColumn: z
			.string()
			.min(1, {
				message: 'Ange namn på artikelnummerkolumn'
			})
			.default('Leverantörens artikelnummer'),
		documentTypeColumn: z
			.string()
			.min(1, {
				message: 'Ange namn på dokumenttypkolumn'
			})
			.default('Dokumenttyp'),
		pdfLinkColumn: z
			.string()
			.min(1, {
				message: 'Ange namn på PDF-länkkolumn'
			})
			.default('Filnamn eller webblänk'),
		company: z.enum(['illbruck', 'nullifire', 'matacryl', 'vandex', 'flowcrete', 'tremco']),
		updateDate: z.boolean().default(false),
		dateColumn: z.string(),
		useDocumentTypeColumn: z.boolean().default(false)
	})
	.refine((data) => data.endNumber > data.startNumber, {
		message: 'Slutnummer måste vara större än startnumret',
		path: ['endNumber']
	});

export type FormSchema = typeof formSchema;
