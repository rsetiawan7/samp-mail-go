#if !defined n_sampmailgo_include
  #define n_sampmailgo_include
#endif

#include a_http
#include strlib

#define SAMPMAILGO_RESPONSE_OK 200
#define SAMPMAILGO_RESPONSE_FORBIDEN 403
#define SAMPMAILGO_RESPONSE_ERROR 400
#define SAMPMAILGO_RESPONSE_NOTFOUND 404

#define SAMPMAILGO_URL "127.0.0.1:8080" // IP + PORT of the server running the NodeJS script
#define SAMPMAILGO_PASSWORD "changeme" // This password should be the same as the NodeJS script. Remove this for disable password
#define SAMPMAILGO_CONSOLELOG 0 // Log some information on console
#define SAMPMAILGO_MAXPOSTDATA 1024 // Cellsize of the variable that holds the data sent to the script. Lower for small emails, higher for more extense emails

new uniqueID = 0;

forward HandleHTTPResponse(index, response_code, data[]);
public HandleHTTPResponse(index, response_code, data[]) {
  uniqueID++;

  #if SAMPMAILGO_CONSOLELOG == 1
  printf("\n[SAMPMAILGO] Server returned a response %d", response_code);
  #endif

  if (response_code == SAMPMAILGO_RESPONSE_OK) {
    #if SAMPMAILGO_CONSOLELOG == 1
    printf("\n[SAMPMAILGO] Email %d sent successfully", uniqueID);
    #endif
  }

	return response_code;
}

stock SendEmail(name[], to[], subject[], text[], bool:isTemplate = false, templateName[] = "") {
	new string[128], postData[SAMPMAILGO_MAXPOSTDATA];

  format(string, sizeof string, "%s", SAMPMAILGO_URL, SAMPMAILGO_PASSWORD);

  #if defined SAMPMAILGO_PASSWORD
  format(string, sizeof string, "%s/?password=%s", string, SAMPMAILGO_PASSWORD);
  #endif

	if (isTemplate)
  {
		format(string, sizeof string, "%s&template=%s", string, templateName);
	}

  format(postData, sizeof postData, "%s|%s|%s|%s", name, to, subject, text);

  utf8encode(postData, postData);
	return HTTP(uniqueID, HTTP_POST, string, postData, "HandleHTTPResponse");
}
