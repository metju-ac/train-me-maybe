export const USER_COOKIE = "user";

export function getCookie(name: string): string | null {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop()?.split(";").shift() || null;

  // TODO REMOVE:
  return "testKU";

  return null;
}
