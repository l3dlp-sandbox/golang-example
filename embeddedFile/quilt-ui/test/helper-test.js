import { expect } from 'chai';
import { parseQUIPrefix, splitHost, splitSubDomain, getEnvironmentAndCountryFromDomain, getDomainFromEnvironmentAndCountry, checkIsLocalEnvironment } from '../src/helper'

describe('#parseQUIPrefix', function() {
    it('should handle null input', function() {
        const {isLocale, parsedKeyValue} = parseQUIPrefix(null);
        expect(isLocale).to.be.false;
        expect(parsedKeyValue).to.eql([{"": {
            value: null,
            editing: false,
        }}]);
    });
    
    it('should handle input with no locales', function() {
        const {isLocale, parsedKeyValue} = parseQUIPrefix("20 sms credits");
        expect(isLocale).to.be.false;
        expect(parsedKeyValue).to.eql([{"":{
            value: "20 sms credits",
            editing: false,
        }}]);
    });

    it('should handle input with single locales', function() {
        const {isLocale, parsedKeyValue} = parseQUIPrefix("20 QUI1SMS_KEYQUI2smsQUI3 credits");
        expect(isLocale).to.be.true;
        expect(parsedKeyValue).to.eql([
            {
                "": {
                    value: "20 ", 
                    editing: false,
                }
            },{
                "SMS_KEY": {
                    value: "sms",
                    editing: false,
                }
            }, {
                "": {
                    value: " credits", 
                    editing: false,
                }
            }
        ]);
    });

    it('should handle input with multiple locales', function() {
        const {isLocale, parsedKeyValue} = parseQUIPrefix("20 QUI1SMS_KEYQUI2smsQUI3 credits QUI1LEFT_KEYQUI2leftQUI3 till EOM");
        expect(isLocale).to.be.true;
        expect(parsedKeyValue).to.eql([
            {
                "": {
                    value: "20 ", 
                    editing: false,
                }
            },{
                "SMS_KEY": {
                    value: "sms",
                    editing: false,
                }
            }, {
                "": {
                    value: " credits ",
                    editing: false,
                }
            },{
                "LEFT_KEY": {
                    value: "left",
                    editing: false,
                },
            }, {
                "": {
                    value: " till EOM", 
                    editing: false,
                }
            }
        ]);
    });

    // Todo, uncomment this test once page editor edit section is implemented and it's clearer how we will need to handle nested prefixes.

    // it('should handle input with nested locales', function() {
    //     const parsedValue = parseQUIPrefix("/nQUI1expires_in_less_than_hoursQUI2Expires in less than 10 QUI1hoursQUI2hoursQUI3.QUI3/n");
    //     expect(parsedValue).to.equal("/nExpires in less than 10 hours./n");
    // });
});

describe('#splitHost', function() {
    it('should split host correctly if host is local', function() {
        const host = "ssg-quilt.circles.local"
        const {subdomain, domain} = splitHost(host);
        expect(subdomain).to.eql("ssg-quilt");
        expect(domain).to.eql("circles.local");
    });
    it('should split host correctly if host is not local', function() {
        const host = "ssg-quilt.circles.life"
        const {subdomain, domain} = splitHost(host);
        expect(subdomain).to.eql("ssg-quilt");
        expect(domain).to.eql("circles.life");
    });
})

describe('#splitSubDomain', function() {
    it('should split subdomain correctly if sg stage', function() {
        const subdomain = "ssg-quilt"
        const {productionString, environmentString} = splitSubDomain(subdomain);
        expect(productionString).to.eql("s");
        expect(environmentString).to.eql("sg");
    });
    it('should split subdomain correctly if id prod', function() {
        const subdomain = "pidgc-quilt"
        const {productionString, environmentString} = splitSubDomain(subdomain);
        expect(productionString).to.eql("p");
        expect(environmentString).to.eql("idgc");
    });
})

describe('#checkIsLocalEnvironment', function() {
    it('should return true if is local environment', function() {
        const host = "ssg-quilt.circles.local"
        expect(checkIsLocalEnvironment(host)).to.be.true
    });
    it('should return false if not local environment', function() {
        const host = "ssg-quilt.circles.life"
        expect(checkIsLocalEnvironment(host)).to.be.false
    });
})

describe('#getDomainFromEnvironmentAndCountry', function() {
    const isLocalEnvironmentOption = [true, false]
    const environmentOptions = ["stage", "preprod", "prod"]
    const countryOptions = ["SG", "TW", "AU", "JP", "ID"]
    for (let isLocalEnvironment of isLocalEnvironmentOption) {
        for (let environmentOption of environmentOptions) {
            for (let countryOption of countryOptions) {
                it(`should return correct domain when isLocalEnvironment is ${isLocalEnvironment} for ${countryOption} ${environmentOption}`, function() {
                    const result = getDomainFromEnvironmentAndCountry(isLocalEnvironment, environmentOption, countryOption);
                    if (isLocalEnvironment) {
                        switch (countryOption) {
                            case "SG": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://ssg-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qsg-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://psg-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "TW": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://stw-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qtw-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://ptw-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "AU": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://sau-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qau-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://pau-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "JP": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://sjp-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qjp-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://pjp-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "ID": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://sid-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qid-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://pidgc-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "JP": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://sjp-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qjp-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://pjp-quilt.circles.local:9991/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                        }
                    } else {
                        switch (countryOption) {
                            case "SG": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("https://ssg-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("https://qsg-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("https://psg-quilt.circles.life/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "TW": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://stw-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qtw-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://ptw-quilt.circles.life/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "AU": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("https://sau-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("https://qau-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("https://pau-quilt.circles.life/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "JP": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("https://sjp-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("https://qjp-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("https://pjp-quilt.circles.life/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "ID": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("http://sid-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("http://qid-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("http://pidgc-quilt.liveon.id/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                            case "JP": {
                                switch (environmentOption) {
                                    case "stage": {
                                        expect(result).to.eql("https://sjp-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "preprod": {
                                        expect(result).to.eql("https://qjp-quilt.circles.life/web/");
                                        break;
                                    }
                                    case "prod": {
                                        expect(result).to.eql("https://pjp-quilt.circles.life/web/");
                                        break;
                                    }
                                }
                                break;
                            }
                        }
                    }
                });
            }
        }
    }
})

describe('#getEnvironmentAndCountryFromDomain', function() {
    const domainSubStringTypes = ["-quilt.circles.life", "-quilt.circles.local:9991"]
    const countryStrings = ["sg", "tw", "au", "jp", "id"]
    const environmentStrings = ["s", "q", "p"]
    for (let domainSubStringType of domainSubStringTypes) {
        for (let environmentString of environmentStrings) {
            for (let countryString of countryStrings) {
                // Handling indonesia which has a different domain pattern for prod
                if (countryString === "id" && environmentString === "p") {
                    const host = `pidgc${domainSubStringType}`
                    it(`should return correct results for ${host}`, function() {
                        const {country, env} = getEnvironmentAndCountryFromDomain(host);
                        expect(country).to.eql("ID");
                        expect(env).to.eql("prod");
                    })
                } else {
                    const host = `${environmentString}${countryString}${domainSubStringType}`
                    it(`should return correct results for ${host}`, function() {
                        const {country, env} = getEnvironmentAndCountryFromDomain(host);
                        switch (countryString) {
                            case "sg": {
                                expect(country).to.eql("SG");
                                break;
                            }
                            case "tw": {
                                expect(country).to.eql("TW");
                                break;
                            }
                            case "au": {
                                expect(country).to.eql("AU");
                                break;
                            }
                            case "jp": {
                                expect(country).to.eql("JP");
                                break;
                            }
                        }
                        switch (environmentString) {
                            case "s": {
                                expect(env).to.eql("stage");
                                break;
                            }
                            case "q": {
                                expect(env).to.eql("preprod");
                                break;
                            }
                            case "p": {
                                expect(env).to.eql("prod");
                                break;
                            }
                        }
                    });
                }
            }
        }
    }
})
