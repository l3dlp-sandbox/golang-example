import { map, get } from 'lodash';
import React from 'react';

import EditorFieldComponent from './components/EditorFieldComponent';
import RequestService from './requestService';

const QUI = "QUI"
const QUI1 = "QUI1"
const QUI2 = "QUI2"
const QUI3 = "QUI3"
const WRITE_ACCESS = "WRITE_ACCESS"

// getCookie retrieves the cookie in the browser based on the given name
// TODO: If there are multiple cookies with the same name this function becomes buggy.
export const getCookie = (name) => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

// This function checks if value exists and if it contains a QUI prefix
// If it contains QUI prefix, it will return the parsed key value object
// Sample of a localization key value is QUI1somekeyQUI2somevalueQUI3
// Example input: "20 QUI1SMS_KEYQUI2smsQUI3 credits QUI1LEFT_KEYQUI2leftQUI3 till EOM" 
// Check out the test cases that covers this function for more details
export const parseQUIPrefix = (value) => {
    // Return if value doesnt contain QUI prefix
    if (!value || !value.includes(QUI)) {
        // isLocale key is used to determine if we need to render an editorField for this value
        return {
            isLocale: false,
            parsedKeyValue: [{"": {
                value,
                editing: false,
            }}],
        }
    }
    // parsedKeyValue array is used to store a collection of objects that can be either a 1. "non locale value", 2. "locale value"
    // This is used to reconstruct the value string for rendering in the preview as well as rendering the editorFieldComponent
    // 1. "non locale value"
    // Object will look like:
    // {"": {
    //     value: "some non locale value",
    //     editing: false,
    // }}
    // 
    // 2. "locale value"
    // Object will look like:
    // {[localeKey]: {
    //     value: localeValue,
    //     editing: false,
    // }}
    const parsedKeyValue = []
    // First split value by the first QUI prefix. 
    // This will give an array like this based on the Example input above: 
    // ["20 ", "SMS_KEYQUI2smsQUI3 credits ", "LEFT_KEYQUI2leftQUI3 till EOM"]
    map(value.split(QUI1), (subStr) => {
        // Check if subStr contains a QUI3 prefix
        if (subStr.includes(QUI3)) {
            // If so, split substring by QUI3 to separate localeKeyValue and noLocaleString.
            // Based on the 2nd substring in the example array above "SMS_KEYQUI2smsQUI3 credits "
            // localeKeyValue will be "SMS_KEYQUI2sms"
            // noLocaleString will be " credits "
            const [localeKeyValue, noLocaleString] = subStr.split(QUI3)
            // Next split substring by QUI2 to get localeKey and localeValue
            // Example "SMS_KEYQUI2sms" will be split to localeKey="SMS_KEY" and localeValue="sms"
            const [localeKey, localeValue] = localeKeyValue.split(QUI2)
            // Push localeKey and localeValue into parsedKeyValue array
            // Instantiate editing to false as this is needed to render in the UI whether the localeValue is being edited
            parsedKeyValue.push({[localeKey]: {
                value: localeValue,
                editing: false,
            }})
            // If noLocaleString present, push into parsedKeyValue array as well. This is needed to reconstruct the value string for rendering in the preview
            if (noLocaleString) {
                parsedKeyValue.push({[""]: {
                    value: noLocaleString,
                    editing: false,
                }})
            }
        } else {
            // If subStr does not contain a QUI prefix, push into parsedKeyValue array as well. Same rationale as noLocaleString above.
            parsedKeyValue.push({[""]: {
                value: subStr,
                editing: false,
            }})
        }
    })
    return {
        isLocale: true,
        parsedKeyValue,
    }
}

// TODO: Add tests
// getValueFromParsedObject is the helper function to reconstruct the value string from the parsedKeyValue key returned from parseQUIPrefix function.
export const getValueFromParsedObject = (parsedObject) => {
    // Example parsedObject:
    // [
    //     {"": {
    //         value: "$20",
    //         editing: false,
    //     }},
    //     {monthly_recurring: {
    //         value: "/mo",
    //         editing: false,
    //     }}
    // ]
    // 
    const parsedValue = map(parsedObject, (obj) => {
        // For each obj, retrieve the inner value key.
        // {monthly_recurring: {
        //     value: ", /mo",
        //     editing: false,
        // }}
        const valueArray = map(Object.values(obj), objectValues => {
            return objectValues.value
        })
        // join valueArray to reconstruct the original value string. 
        // From the example above it will look like: "/mo"
        return valueArray.join("")
    }).join("")
    // join the result of the map:
    // From the example above it will look like: "$20/mo"
    return parsedValue
}

// TODO: Add tests
// getValueFromWidgetData first retrieves the parsedKeyValue from each of the widget component keys
// It then uses the getValueFromParsedObject function to reconstruction it's original value.
export const getValueFromWidgetData = (widgetData, type, key) => {
    // Example type: addon-tile
    // Example key: data.title
    // Example widgetData:
    // {
    //     "addon-tile": {
    //         "data.color": {
    //             "isLocale": false,
    //             "parsedKeyValue": [
    //                 {
    //                     "": {
    //                         "value": "#B766EA",
    //                         "editing": false
    //                     }
    //                 }
    //             ]
    //         },
    //         "data.title": {
    //             "isLocale": true,
    //             "parsedKeyValue": [
    //                 {
    //                     "": {
    //                         "value": "$20",
    //                         "editing": false
    //                     },
    //                     "monthly_recurring": {
    //                         "value": ""/mo",
    //                         "editing": false
    //                     }
    //                 }
    //             ]
    //         },
    //     }
    // }
    const parsedKeyValue = get(widgetData, [type, key, "parsedKeyValue"])
    // Based on example widgetData input above for type: "addon-tile" and key: "data.title", parsedKeyValue would be:
    // [
    //     {
    //         "": {
    //             "value": "$20",
    //             "editing": false
    //         }
    //     },
    //     {
    //         "monthly_recurring": {
    //             "value": "/mo",
    //             "editing": false
    //         }
    //     }
    // ]
    const parsedValue = getValueFromParsedObject(parsedKeyValue)
    return parsedValue
}

// setKeyValueToState is used to update the widget's widgetData state with the parsedKeyValue object.
export const setKeyValueToState = (data, type, keyName, updateStateFunction) => {
    const { isLocale, parsedKeyValue } = parseQUIPrefix(get(data, keyName))
    updateStateFunction(prevState => {
        const newValue = {...prevState,
            [type]: {
                ...prevState[type],
                [keyName]: {
                    isLocale,
                    parsedKeyValue,
                }
            }
        }
        return newValue
    })
}

// renderEditorFieldComponents renders the EditorFieldComponent based on the widgetData that was configured by setKeyValueToState
export const renderEditorFieldComponents = ({widgetData, saving, setWidgetData, editing, setEditing, initState, updateLocalisationValueFunc}) => {
    const editorFieldComponents = []
    const fieldCount = Object.keys(widgetData)
    if (fieldCount != 0) {
        map(widgetData, (typeObject, typeKey) => {
            map(typeObject, (valueObject, valueKey) => {
                // Only render EditorFieldComponent for values that are editable locales
                if (valueObject.isLocale) {
                    const parsedKeyValueCount = get(valueObject, "parsedKeyValue.length", 0)
                    if (parsedKeyValueCount > 0) {
                        // Loop through all the parsedKeyValue objects
                        for(let index=0; index < parsedKeyValueCount; index+=1) {
                            let localeObject = valueObject.parsedKeyValue[index]
                            let localeKey = Object.keys(localeObject)[0]
                            // Render EditorFieldComponent for parsedKeyValue objects that are editable locales
                            if (localeKey != "") {
                                editorFieldComponents.push(
                                    <EditorFieldComponent
                                        key={`${typeKey}-${valueKey}-${localeKey}`}
                                        localeObject={localeObject}
                                        localeKey={localeKey}
                                        editing={editing}
                                        setEditing={setEditing}
                                        widgetData={widgetData}
                                        setWidgetData={setWidgetData}
                                        initState={initState}
                                        saving={saving}
                                        typeKey={typeKey}
                                        valueKey={valueKey}
                                        index={index}
                                        updateLocalisationValueFunc={updateLocalisationValueFunc}
                                    />
                                )
                            }
                        }
                    }
                }
            })
        })
    }
    return editorFieldComponents
}

// updateLocalisationValue invokes the Quilt API to update the localeValue for the selected localeKey
export const updateLocalisationValue = ({localeKey, localeValue, pageId, setErrorMessage, setSaving, setEditing, fetchPageData}) => {
    (async () => {
        const accessType = getCookie("X-ACCESS-TYPE")
        if (accessType !== WRITE_ACCESS) {
            setErrorMessage(`Your access type ${accessType} doesn't allow you to update the locale`)
            return
        }
        console.log(`Updating: ${localeKey}, ${localeValue}`)
        setSaving(true)
        try {
            const res = await RequestService.put(`${process.env.QUILT_DOMAIN}api/v1/quilt/ui/page/${pageId}`, { 
                updateKey: localeKey, 
                updateValue: localeValue 
            }, {})
            if (res.status === 200) {
                setEditing(false)
                fetchPageData()
            } else {
                console.log("Error updating localisation value", res)
            }
        } catch (error) {
            if (error.response && error.response.data && error.response.data.description) {
                setErrorMessage(error.response.data.description)
            } else {
                setErrorMessage("Error updating locale")
            }
        }
        setSaving(false)
    })()
}

export const splitHost = (host) => {
    const splitResult = host.split(".")
    const subdomain = splitResult[0]
    const domain = splitResult.slice(-2).join(".")
    return { subdomain, domain }
}

export const checkIsLocalEnvironment = (host) => host.includes("circles.local") ? true : false

export const splitSubDomain = (subdomain) => {
    const [countryEnv] = subdomain.split("-quilt")
    const productionString = countryEnv[0]
    const environmentString = countryEnv.slice(1)
    return { productionString, environmentString }
}

export const getDomainFromEnvironmentAndCountry = (isLocalEnvironment, environment, country) => {
    if (isLocalEnvironment) {
        switch (environment) {
            case "stage":
                switch (country) {
                    case "SG":
                        return "http://ssg-quilt.circles.local:9991/web/"
                    case "TW":
                        return "http://stw-quilt.circles.local:9991/web/"
                    case "AU":
                        return "http://sau-quilt.circles.local:9991/web/"
                    case "ID":
                        return "http://sid-quilt.circles.local:9991/web/"
                    case "JP":
                        return "http://sjp-quilt.circles.local:9991/web/"
                }
            case "preprod":
                switch (country) {
                    case "SG":
                        return "http://qsg-quilt.circles.local:9991/web/"
                    case "TW":
                        return "http://qtw-quilt.circles.local:9991/web/"
                    case "AU":
                        return "http://qau-quilt.circles.local:9991/web/"
                    case "ID":
                        return "http://qid-quilt.circles.local:9991/web/"
                    case "JP":
                        return "http://qjp-quilt.circles.local:9991/web/"
                }
            case "prod":
                switch (country) {
                    case "SG":
                        return "http://psg-quilt.circles.local:9991/web/"
                    case "TW":
                        return "http://ptw-quilt.circles.local:9991/web/"
                    case "AU":
                        return "http://pau-quilt.circles.local:9991/web/"
                    case "ID":
                        return "http://pidgc-quilt.circles.local:9991/web/"
                    case "JP":
                        return "http://pjp-quilt.circles.local:9991/web/"
                }
        }
    } else {
        switch (environment) {
            case "stage":
                switch (country) {
                    case "SG":
                        return "https://ssg-quilt.circles.life/web/"
                    case "TW":
                        return "http://stw-quilt.circles.life/web/"
                    case "AU":
                        return "https://sau-quilt.circles.life/web/"
                    case "ID":
                        return "http://sid-quilt.circles.life/web/"
                    case "JP":
                        return "https://sjp-quilt.circles.life/web/"
                }
            case "preprod":
                switch (country) {
                    case "SG":
                        return "https://qsg-quilt.circles.life/web/"
                    case "TW":
                        return "http://qtw-quilt.circles.life/web/"
                    case "AU":
                        return "https://qau-quilt.circles.life/web/"
                    case "ID":
                        return "http://qid-quilt.circles.life/web/"
                    case "JP":
                        return "https://qjp-quilt.circles.life/web/"
                }
            case "prod":
                switch (country) {
                    case "SG":
                        return "https://psg-quilt.circles.life/web/"
                    case "TW":
                        return "http://ptw-quilt.circles.life/web/"
                    case "AU":
                        return "https://pau-quilt.circles.life/web/"
                    case "ID":
                        return "http://pidgc-quilt.liveon.id/web/"
                    case "JP":
                        return "https://pjp-quilt.circles.life/web/"
                }
        }
    }
}

export const getEnvironmentAndCountryFromDomain = (host) => {
    const { subdomain, domain } = splitHost(host)
    const { productionString, environmentString } = splitSubDomain(subdomain)

    let env
    let country
    switch (productionString) {
        case "s":
            env = "stage"
            break;
        case "q":
            env = "preprod"
            break;
        case "p":
            env = "prod"
            break;
        default:
            env = "stage"
    }

    switch (environmentString) {
        case "sg":
            country = "SG"
            break;
        case "tw":
            country = "TW"
            break;
        case "au":
            country = "AU"
            break;
        case "id":
        case "idgc":
            country = "ID"
            break;
        case "jp":
            country = "JP"
            break;
        default:
            country = "SG"
    }
    return { country, env }
}
