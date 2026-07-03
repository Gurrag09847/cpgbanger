import "clsx";
import { a0 as attr_class, a1 as bind_props, a2 as ensure_array_like, a3 as store_get, a4 as unsubscribe_stores } from "../../chunks/index2.js";
import { c as clsx, b as attr, e as escape_html } from "../../chunks/attributes.js";
import { w as writable } from "../../chunks/index.js";
function FileUpload($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    $$renderer2.push(`<div${attr_class(`border-border hover:border-primary/50 bg-muted cursor-pointer rounded-lg border-2 border-dashed p-8 text-center transition-colors ${""} ${""}`)} role="button" tabindex="0">`);
    {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`<svg class="text-muted-foreground mx-auto mb-3 h-10 w-10" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path></svg> <p class="text-foreground font-medium">Släpp din .xlsx-fil här</p> <p class="text-muted-foreground mt-1 text-sm">eller klicka för att bläddra</p>`);
    }
    $$renderer2.push(`<!--]--></div> <input id="file-input" type="file" accept=".xlsx" class="hidden"/> `);
    {
      $$renderer2.push("<!--[-1-->");
    }
    $$renderer2.push(`<!--]-->`);
  });
}
function FileSelector($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let { activeFileId = null, activeFileName = null, onRefresh } = $$props;
    {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-muted-foreground text-sm">Laddar filer...</p>`);
    }
    $$renderer2.push(`<!--]-->`);
    bind_props($$props, { activeFileId, activeFileName });
  });
}
function ConfigForm($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let { activeFileId } = $$props;
    const companies = [
      { value: "illbruck", label: "illbruck" },
      { value: "vandex", label: "vandex" },
      { value: "nullifire", label: "nullifire" },
      { value: "flowcrete", label: "flowcrete" },
      { value: "matacryl", label: "matacryl" },
      { value: "tremco", label: "tremco" }
    ];
    const docTypes = [
      { value: "", label: "-- Välj --" },
      { value: "Produktdatablad", label: "Produktdatablad" },
      { value: "Säkerhetsdatablad", label: "Säkerhetsdatablad" },
      { value: "Prestandadeklaration", label: "Prestandadeklaration" },
      {
        value: "Miljövarudeklaration (EPD)",
        label: "Miljövarudeklaration (EPD)"
      },
      { value: "Certifikat", label: "Certifikat" }
    ];
    let form = {
      sheetName: "",
      sheetNamesStr: "",
      artikelnummerCol: "",
      pdfLinkCol: "",
      startNumber: 1,
      endNumber: 100,
      company: "illbruck",
      documentType: "",
      useDocumentTypeColumn: false,
      updateDate: false,
      concurrency: 5
    };
    const inputClass = "border-border bg-background placeholder:text-muted-foreground text-foreground w-full rounded-md border px-3 py-2 text-sm outline-none focus:border-primary focus:ring-1 focus:ring-primary";
    const labelClass = "text-foreground mb-1 block text-sm font-medium";
    const checkboxClass = "border-border bg-background text-primary h-4 w-4 rounded border focus:ring-primary";
    $$renderer2.push(`<form class="space-y-4"><div class="grid gap-4 sm:grid-cols-2"><div><label${attr_class(clsx(labelClass))} for="sheetName">Bladnamn</label> <input id="sheetName"${attr_class(clsx(inputClass))} type="text"${attr("value", form.sheetName)} placeholder="Blad1"/></div> <div><label${attr_class(clsx(labelClass))} for="sheetNamesStr">Flera blad (kommaseparerat)</label> <input id="sheetNamesStr"${attr_class(clsx(inputClass))} type="text"${attr("value", form.sheetNamesStr)} placeholder="t.ex. Blad1, Blad2"/></div> <div><label${attr_class(clsx(labelClass))} for="company">Företag</label> `);
    $$renderer2.select({ id: "company", class: inputClass, value: form.company }, ($$renderer3) => {
      $$renderer3.push(`<!--[-->`);
      const each_array = ensure_array_like(companies);
      for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
        let c = each_array[$$index];
        $$renderer3.option({ value: c.value }, ($$renderer4) => {
          $$renderer4.push(`${escape_html(c.label)}`);
        });
      }
      $$renderer3.push(`<!--]-->`);
    });
    $$renderer2.push(`</div> <div><label${attr_class(clsx(labelClass))} for="artikelnummerCol">Artikelnummer-kolumn</label> <input id="artikelnummerCol"${attr_class(clsx(inputClass))} type="text"${attr("value", form.artikelnummerCol)} placeholder="t.ex. Artikelnummer" required=""/></div> <div><label${attr_class(clsx(labelClass))} for="pdfLinkCol">PDF-länk-kolumn</label> <input id="pdfLinkCol"${attr_class(clsx(inputClass))} type="text"${attr("value", form.pdfLinkCol)} placeholder="t.ex. Filnamn eller webblänk"/></div> <div><label${attr_class(clsx(labelClass))} for="startNumber">Startrad</label> <input id="startNumber"${attr_class(clsx(inputClass))} type="number" min="1"${attr("value", form.startNumber)} required=""/></div> <div><label${attr_class(clsx(labelClass))} for="endNumber">Slutrad</label> <input id="endNumber"${attr_class(clsx(inputClass))} type="number" min="1"${attr("value", form.endNumber)} required=""/></div> <div><label${attr_class(clsx(labelClass))} for="concurrency">Samtidiga anrop</label> <input id="concurrency"${attr_class(clsx(inputClass))} type="number" min="1" max="20"${attr("value", form.concurrency)}/></div></div> <div class="border-border border-t pt-4"><label class="flex items-center gap-2 text-sm"><input type="checkbox"${attr_class(clsx(checkboxClass))}${attr("checked", form.useDocumentTypeColumn, true)}/> <span class="text-foreground">Använd dokumenttypskolumn från kalkylblad</span></label> `);
    {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`<div class="mt-3"><label${attr_class(clsx(labelClass))} for="docType">Dokumenttyp</label> `);
      $$renderer2.select({ id: "docType", class: inputClass, value: form.documentType }, ($$renderer3) => {
        $$renderer3.push(`<!--[-->`);
        const each_array_1 = ensure_array_like(docTypes);
        for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
          let dt = each_array_1[$$index_1];
          $$renderer3.option({ value: dt.value }, ($$renderer4) => {
            $$renderer4.push(`${escape_html(dt.label)}`);
          });
        }
        $$renderer3.push(`<!--]-->`);
      });
      $$renderer2.push(`</div>`);
    }
    $$renderer2.push(`<!--]--></div> <div class="border-border border-t pt-4"><label class="flex items-center gap-2 text-sm"><input type="checkbox"${attr_class(clsx(checkboxClass))}${attr("checked", form.updateDate, true)}/> <span class="text-foreground">Uppdatera datumkolumn på varje rad</span></label> `);
    {
      $$renderer2.push("<!--[-1-->");
    }
    $$renderer2.push(`<!--]--></div> `);
    {
      $$renderer2.push("<!--[-1-->");
    }
    $$renderer2.push(`<!--]--> <button type="submit" class="bg-primary text-primary-foreground hover:bg-primary/90 inline-flex w-full items-center justify-center rounded-md px-4 py-2.5 text-sm font-medium transition-colors disabled:opacity-50"${attr("disabled", !activeFileId, true)}>`);
    {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`Kör nu`);
    }
    $$renderer2.push(`<!--]--></button> `);
    if (!activeFileId) {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-muted-foreground text-xs text-center">Välj en aktiv fil ovan för att aktivera kör-knappen</p>`);
    } else {
      $$renderer2.push("<!--[-1-->");
    }
    $$renderer2.push(`<!--]--></form>`);
  });
}
function SchedulePanel($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-muted-foreground text-sm">Laddar schema...</p>`);
    }
    $$renderer2.push(`<!--]-->`);
  });
}
function JobPanel($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-muted-foreground text-sm">Laddar jobb...</p>`);
    }
    $$renderer2.push(`<!--]-->`);
  });
}
const toasts = writable([]);
toasts.subscribe((v) => {
});
function Toaster($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    var $$store_subs;
    $$renderer2.push(`<div class="pointer-events-none fixed right-4 top-4 z-50 flex flex-col gap-2"><!--[-->`);
    const each_array = ensure_array_like(store_get($$store_subs ??= {}, "$toasts", toasts));
    for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
      let toast = each_array[$$index];
      $$renderer2.push(`<button${attr_class(`pointer-events-auto border-border bg-card text-foreground max-w-sm rounded-lg border px-4 py-3 text-left text-sm shadow-lg ${toast.variant === "error" ? "border-red-500" : ""} ${toast.variant === "success" ? "border-green-500" : ""}`)}>${escape_html(toast.message)}</button>`);
    }
    $$renderer2.push(`<!--]--></div>`);
    if ($$store_subs) unsubscribe_stores($$store_subs);
  });
}
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let activeFileId = null;
    let activeFileName = null;
    let refreshJobs = 0;
    let $$settled = true;
    let $$inner_renderer;
    function $$render_inner($$renderer3) {
      $$renderer3.push(`<div class="space-y-8"><div class="grid gap-8 lg:grid-cols-2"><section class="space-y-6"><div class="border-border bg-card rounded-lg border p-6"><h2 class="text-foreground mb-4 text-lg font-semibold">1. Ladda upp Excel-fil</h2> `);
      FileUpload($$renderer3);
      $$renderer3.push(`<!----></div> <div class="border-border bg-card rounded-lg border p-6"><h2 class="text-foreground mb-4 text-lg font-semibold">2. Välj fil</h2> `);
      FileSelector($$renderer3, {
        onRefresh: () => refreshJobs++,
        get activeFileId() {
          return activeFileId;
        },
        set activeFileId($$value) {
          activeFileId = $$value;
          $$settled = false;
        },
        get activeFileName() {
          return activeFileName;
        },
        set activeFileName($$value) {
          activeFileName = $$value;
          $$settled = false;
        }
      });
      $$renderer3.push(`<!----></div></section> <section class="space-y-6"><div class="border-border bg-card rounded-lg border p-6"><h2 class="text-foreground mb-4 text-lg font-semibold">3. Hämtningsinställningar</h2> `);
      ConfigForm($$renderer3, {
        activeFileId
      });
      $$renderer3.push(`<!----></div> <div class="border-border bg-card rounded-lg border p-6"><h2 class="text-foreground mb-4 text-lg font-semibold">4. Schemalägg</h2> `);
      SchedulePanel($$renderer3);
      $$renderer3.push(`<!----></div></section></div> <div class="border-border bg-card rounded-lg border p-6"><h2 class="text-foreground mb-4 text-lg font-semibold">Jobbhistorik</h2> `);
      JobPanel($$renderer3);
      $$renderer3.push(`<!----></div></div> `);
      Toaster($$renderer3);
      $$renderer3.push(`<!---->`);
    }
    do {
      $$settled = true;
      $$inner_renderer = $$renderer2.copy();
      $$render_inner($$inner_renderer);
    } while (!$$settled);
    $$renderer2.subsume($$inner_renderer);
  });
}
export {
  _page as default
};
