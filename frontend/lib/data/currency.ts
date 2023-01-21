export type Codes = "CHF" | "USD" | "EUR" | "GBP" | "JPY" | "ZAR" | "AUD" | "NOK" | "SEK" | "DKK" | "RMB";

export type Currency = {
  code: Codes;
  local: string;
  symbol: string;
  name: string;
};

export const currencies: Currency[] = [
  {
    code: "CHF",
    local: "de-CH",
    symbol: "CHF",
    name: "Swiss Franc",
  },
  {
    code: "AUD",
    local: "en-AU",
    symbol: "$",
    name: "Australian Dollar",
  },
  {
    code: "GBP",
    local: "en-GB",
    symbol: "£",
    name: "British Pound",
  },
  {
    code: "RMB",
    local: "zh-CN",
    symbol: "¥",
    name: "Chinese Yuan",
  },
  {
    code: "DKK",
    local: "da-DK",
    symbol: "kr",
    name: "Danish Krone",
  },
  {
    code: "EUR",
    local: "de-DE",
    symbol: "€",
    name: "Euro",
  },
  {
    code: "JPY",
    local: "ja-JP",
    symbol: "¥",
    name: "Japanese Yen",
  },
  {
    code: "NOK",
    local: "nb-NO",
    symbol: "kr",
    name: "Norwegian Krone",
  },
  {
    code: "ZAR",
    local: "en-ZA",
    symbol: "R",
    name: "South African Rand",
  },
  {
    code: "SEK",
    local: "sv-SE",
    symbol: "kr",
    name: "Swedish Krona",
  },
  {
    code: "USD",
    local: "en-US",
    symbol: "$",
    name: "US Dollar",
  },
];
