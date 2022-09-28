import React, { useState, useEffect } from 'react';
import { useLocation, useHistory } from 'react-router-dom';
import {
  Grid,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  makeStyles,
} from '@material-ui/core';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import { map } from 'lodash';

import PageEditor from './PageEditor'; 
import PageLayout from '../components/PageLayout';
import { getEnvironmentAndCountryFromDomain, getDomainFromEnvironmentAndCountry, checkIsLocalEnvironment } from '../helper';

export const COUNTRY_LOCALE = (country) => {
    let locales = {}; 

    switch (country) {
        case "SG":
        case "AU": {
            locales = {
                "English": "en"
            }
            break;
        }
        case "JP": {
            locales = {
                "Japanese": "ja",
                "English": "en"
            }
            break;
        }
        case "ID": {
            locales = {
                "English": "en",
                "Indonesian": "id"
            }
            break;
        }
        case "TW": {
            locales = {
                "English": "en",
                "Chinese": "zh"
            }
            break;
        }
        default: {
            locales = []
        }
    }
    return locales
}

const COUNTRIES = {
    "singapore": "SG",
    "Australia": "AU",
    "Taiwan": "TW",
    "Indonesia": "ID",
    "Japan": "JP",
}

const ENVIRONMENTS = {
    "staging": "stage",
    "pre-production": "preprod",
    "production": "prod",
}

const useStyle = makeStyles(() => ({
  root: {
    flexGrow: 1,
    flexShrink: 10, 
  },
  icon: { color: 'blue' },
}));

const SUPPORTED_PAGES = ['dashboard-ui-test', 'app-loyalty-v2-ui-test', 'dashboard', 'app-loyalty-v2']

const getPageIdFromURL = () => {
    // Get pageId from url path
    const location = window.location.pathname;
    return location.replace('/web/', '').split("?")[0]; 
    // const pageId = location.replace('/web/', '').split("?")[0]
    // return pageId
}

function useQuery() {
    return new URLSearchParams(useLocation().search);
}

const getDefaultLocale = (currentCountry) => {
    const localesInCountry = COUNTRY_LOCALE(currentCountry)
    const localeValuesInCountry = Object.values(localesInCountry)
    return localeValuesInCountry[0]
}

export default function DefaultDropDown() {
    const classes = useStyle();
    const history = useHistory();
    const query = useQuery();

    // data structure of dropdowns using UseState
    const [selectedPage, setSelectedPage] = useState("")
    const [selectedLocale, setSelectedLocale] = useState("")
    const [selectedCountry, setSelectedCountry] = useState("")
    const [selectedEnvironment, setSelectedEnvironment] = useState("")
    const [dropdownValue, setValue] = useState({
        environment: 'staging',
        country: 'SG',
        locale: '',
        pages: '', 
        plan_type: '', 
    });
    // Throw error if user using localhost
    const { host } = window.location
    if (host.includes("localhost")) {
        throw new Error("Quilt UI doesnt work with localhost, checkout readme to configure /etc/hosts for proper behaviour")
    }
    
    const initEnvAndCountry = () => {
        const { country, env } = getEnvironmentAndCountryFromDomain(host)
        setSelectedCountry(country)
        setSelectedEnvironment(env)
    }

    // Whenever country or environment changes and there is a pageID in url...
    // First check if the pageID matches supported list of pages.
        // If default page, this means that no pageID is set. Do not render PageEditor component
        // If it matches, set selectedPage to the pageID
        // If it doesnt, update url to /default
    // Set page
    const initPage = (country) => {
        const pageIDfromURL = getPageIdFromURL()
        console.log("pageIDfromURL", pageIDfromURL)
        if (pageIDfromURL) {
            if (pageIDfromURL === "default") {
                console.log("Default page", pageIDfromURL)
            }
            else if (SUPPORTED_PAGES.includes(pageIDfromURL)) {
                console.log("Supported page", pageIDfromURL)
                setSelectedPage(pageIDfromURL)
            }
        } else {
            console.log("Not supported page", pageIDfromURL, "redirecting to root page")
            const defaultLocale = getDefaultLocale(country)
            console.log("defaultLocale", defaultLocale)
            history.push(`/default?locale=${defaultLocale}`)
            setSelectedPage("")
        }
    }

    // Whenever country or environment changes check if locale is present in url
    // If it is present, check if it is supported based on the current country
        // If it matches, set selectedLocale
        // If it doesnt, update url to the default country's locale
    const initLocale = (country) => {
        const queryLocale = query.get("locale")
        console.log("queryLocale", queryLocale)
        console.log("country", country)
        const localesInCountry = COUNTRY_LOCALE(country)
        const localeValuesInCountry = Object.values(localesInCountry)
        if (queryLocale && localeValuesInCountry.includes(queryLocale)) {
            console.log("Supported locale", queryLocale)
            setSelectedLocale(queryLocale)
        } else {
            console.log("Not supported locale", queryLocale, "for country", country)
            getDefaultLocale(country)
            const pageIDfromURL = getPageIdFromURL()
            const defaultLocale = getDefaultLocale(country)
            history.push(`/${pageIDfromURL}?locale=${defaultLocale}`)
            setSelectedLocale(defaultLocale)
        }
    }

    useEffect(async () => {
        initEnvAndCountry()
    }, [])
    
    useEffect(async () => {
        if (!selectedCountry || !selectedEnvironment) return
        initPage(selectedCountry)
        initLocale(selectedCountry)
    }, [selectedCountry, selectedEnvironment])

    // When user selects a page, change url to the selected page and update selectedPage
    const updatePage = (pageID) => {
        const queryLocale = query.get("locale")
        history.push(`/${pageID}?locale=${queryLocale}`)
        setSelectedPage(pageID)
    }

    // When user selects a locale, change url to the selected locale and update selectedLocale
    const updateLocale = (locale) => {
        const pageIDfromURL = getPageIdFromURL()
        history.push(`/${pageIDfromURL}?locale=${locale}`)
        setSelectedLocale(locale)
    }

    // When user selects an environment redirect user to correct domain
    const updateEnvironment = (environment) => {
        const isLocalEnvironment = checkIsLocalEnvironment(host)
        const urltoRedirect = getDomainFromEnvironmentAndCountry(isLocalEnvironment, environment, selectedCountry)
        console.log("urltoRedirect", urltoRedirect)
        window.location.href = urltoRedirect
    }

    // When user selects an environment redirect user to correct domain
    const updateCountry = (country) => {
        const isLocalEnvironment = checkIsLocalEnvironment(host)
        const urltoRedirect = getDomainFromEnvironmentAndCountry(isLocalEnvironment, selectedEnvironment, country)
        console.log("urltoRedirect", urltoRedirect)
        window.location.href = urltoRedirect
    }    

    const renderPageEditor = () => {
        if (!selectedPage) return null
        console.log("rendering", selectedPage, selectedLocale, selectedCountry)
        return <PageEditor 
            pageId={selectedPage}
            selectedLocale={selectedLocale}
            country={selectedCountry}
        />
    }

    if (!selectedCountry) return null
    return (
        <PageLayout title="Edit Pages">
            <div className={classes.root}>
                <Grid container direction="row" spacing={3}>
                    <Grid item xs={4}>
                        <FormControl fullWidth>
                            <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                Environment
                            </InputLabel>
                            <Select
                                classes={{ icon: classes.icon }}
                                value={selectedEnvironment}
                                onChange={(e) => {
                                    const val = e.target.value;
                                    updateEnvironment(val);
                                }}
                                IconComponent={ExpandMoreIcon}
                            >
                                {map(ENVIRONMENTS, (envValue, envKey) => (
                                    <MenuItem key={envKey} value={envValue}>
                                        {envKey}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={4}>
                        <FormControl fullWidth>
                            <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                Country
                            </InputLabel>

                            <Select
                                classes={{ icon: classes.icon }}
                                value={selectedCountry}
                                onChange={(e) => {
                                    const val = e.target.value;
                                    updateCountry(val);
                                }}
                                IconComponent={ExpandMoreIcon}
                            >
                                {map(COUNTRIES, (countryValue, countryKey) => (
                                    <MenuItem key={countryKey} value={countryValue}>
                                        {countryKey}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={4}>
                        <FormControl fullWidth>
                            <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                Locale
                            </InputLabel>

                            <Select
                                classes={{ icon: classes.icon }}
                                value={selectedLocale}
                                onChange={(e) => {
                                    const val = e.target.value;
                                    updateLocale(val);
                                }}
                                IconComponent={ExpandMoreIcon}
                            >
                                {map(COUNTRY_LOCALE(selectedCountry), (localeValue, localeKey) => (
                                    <MenuItem key={localeKey} value={localeValue}>
                                        {localeKey}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={4}>
                        <FormControl fullWidth>
                            <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                            Pages
                            </InputLabel>
                            <Select
                                classes={{ icon: classes.icon }}
                                value={selectedPage}
                                onChange={(e) => {
                                    const val = e.target.value;
                                    updatePage(val);
                                }}
                                IconComponent={ExpandMoreIcon}
                            >
                                {map(SUPPORTED_PAGES, (page) => (
                                    <MenuItem key={page} value={page}>
                                        {page}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    
                    <Grid item xs ={4}>
                        <FormControl fullWidth disabled>
                            <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                Plan Type
                            </InputLabel>
                            <Select
                                classes={{ icon: classes.icon }}
                                value={dropdownValue.plan_type}
                                onChange={(e) => {
                                const val = e.target.value;
                                    setValue((prevState) => ({ ...prevState, plan_type: val }));
                                }}
                                IconComponent={ExpandMoreIcon}
                            >
                                <MenuItem value="">
                                    <em>None</em>
                                </MenuItem>

                                {/* use loop to create menuItem  */}
                                <MenuItem value={10}>Ten</MenuItem>
                                <MenuItem value={20}>Twenty</MenuItem>
                                <MenuItem value={30}>Thirty</MenuItem>
                            </Select>
                        </FormControl>
                    </Grid>
                </Grid>

                <Grid container>
                    {renderPageEditor()}
                </Grid>
            </div>
        </PageLayout>
    )
}

