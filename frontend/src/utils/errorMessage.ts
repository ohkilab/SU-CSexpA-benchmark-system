export const sanitizeErrorMessage = (message: string): string => {
  if (!message) return "";

  let sanitized = message;

  sanitized = sanitized.replace(
    /(https?:\/\/[^\s"'`<>?#]+)\?[^\s"'`<>]*/g,
    "$1",
  );

  sanitized = sanitized.replace(/(query is not found:\s*)(.+)$/i, "$1[redacted]");
  sanitized = sanitized.replace(/(query:\s*)(.+)$/i, "$1[redacted]");

  return sanitized;
};