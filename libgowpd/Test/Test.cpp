#include "stdafx.h"
#include "libgowpd.h"

int main()
{
	HRESULT hr;

	hr = CoInitializeEx(NULL, COINIT_MULTITHREADED);
	if (FAILED(hr)) {
		printf("init err %x\n", hr);
	}
	IPortableDeviceManager *pPortableDeviceManager;
	hr = CoCreateInstance(
		CLSID_PortableDeviceManager,
		NULL,
		CLSCTX_INPROC_SERVER,
		IID_IPortableDeviceManager,
		(LPVOID*)&pPortableDeviceManager);
	if (SUCCEEDED(hr)) {
		PnPDeviceID *pPnPDeviceIDs;
		DWORD cPnPDeviceIDs;
		hr = portableDeviceManager_GetDevices(pPortableDeviceManager, &pPnPDeviceIDs, &cPnPDeviceIDs);
		if (SUCCEEDED(hr)) {
			printf("%d devices has been found.\n", cPnPDeviceIDs);

			for (int i = 0; i < cPnPDeviceIDs; i++) {
				PnPDeviceID id = pPnPDeviceIDs[i];

				PWSTR friendlyName;
				PWSTR manufacturer;
				PWSTR description;
				DWORD cFriendlyName = 0;
				DWORD cManufacturer = 0;
				DWORD cDescription = 0;

				hr = portableDeviceManager_GetDeviceFriendlyName(pPortableDeviceManager, id, &friendlyName, &cFriendlyName);
				if (SUCCEEDED(hr)) {
					printf("%ws\n", friendlyName);
					free(friendlyName);
					friendlyName = NULL;
				} else {
					printf("friendlyname wtf 0x%x\n", hr);
				}
				hr = portableDeviceManager_GetDeviceManufacturer(pPortableDeviceManager, id, &manufacturer, &cManufacturer);
				if (SUCCEEDED(hr)) {
					printf("%ws\n", manufacturer);
					free(manufacturer);
					manufacturer = NULL;
				} else {
					printf("manufacturer WTF 0x%x\n", hr);
				}
				hr = portableDeviceManager_GetDeviceDescription(pPortableDeviceManager, id, &description, &cDescription);
				if (SUCCEEDED(hr)) {
					printf("%ws\n", description);
					free(description);
					description = NULL;
				} else {
					printf("description WTF 0x%x\n", hr);
				}

				CoTaskMemFree(id);
			}

			free(pPnPDeviceIDs);
		}
		else {
			printf("get devices WTF %x\n", hr);
		}
	}
	else {
		printf("create manager WTF %x\n", hr);
	}

	PROPVARIANT prop = {0};
	PropVariantInit(&prop);

	prop.vt = VT_LPWSTR;
	prop.wReserved1 = 2;
	prop.wReserved2 = 3;
	prop.wReserved3 = 4;
	prop.pwszVal = (LPWSTR)CoTaskMemAlloc(sizeof(WCHAR) * 5);
	char* testStr = "TEST";
	for (int i = 0; i < 4; i++) {
		prop.pwszVal[i] = testStr[i];
	}
	prop.pwszVal[4] = 0;

	printf("prop.vt:      0x%x\n", prop.vt);
	printf("prop.pwszVal: 0x%016p\n", prop.pwszVal);
	byte* bs = (byte*) &prop;
	for (int i = 0; i < sizeof(PROPVARIANT); i++) {
		printf("0x%x\n", bs[i]);
	}
	printf("LPWSTR size:  0x%x\n", sizeof(LPWSTR));

	PropVariantClear(&prop);

    return 0;
}

