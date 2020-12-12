import { LocalDate, ZonedDateTime } from "js-joda";


export const parseLocalDate = (dateStr: string): LocalDate => {
  return dateStr ? ZonedDateTime.parse(dateStr).toLocalDate() : null;
}

