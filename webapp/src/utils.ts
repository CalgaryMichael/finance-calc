import { DateTimeFormatter, LocalDate, ZonedDateTime } from "js-joda";

const shortDateFormatter = DateTimeFormatter.ofPattern("yyyy-MM");

export const parseLocalDate = (dateStr: string): LocalDate => {
  return dateStr ? ZonedDateTime.parse(dateStr).toLocalDate() : null;
}

export const shortDate = (d: LocalDate): string => {
  return d.format(shortDateFormatter);
};

