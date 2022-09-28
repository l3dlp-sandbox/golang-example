import React from 'react';
import clsx from 'clsx';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import CircularProgress from '@material-ui/core/CircularProgress';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';

import { typography, colors } from '../styles/globalStyles'
import CheckIcon from '../assets/Check.png'
import DiscardIcon from '../assets/Discard.png'
import EditIcon from '../assets/Edit.png'

const useStyles = makeStyles(() => ({
    hide: {
        display: "none"
    },
    editorFieldContainer: {
        marginTop: "1rem",
        marginBottom: "1rem",
        display: "grid",
        height: "3rem",
        gridTemplateColumns: "repeat(7, 1fr)",
        gridTemplateRows: "1rem 2rem",
        gridColumnGap: "0px",
        gridRowGap: "0px",
    },
    textField: {
        gridColumnStart: "1",
        gridColumnEnd: "8",
        gridRowStart: "1",
        gridRowEnd: "2",
        marginBottom: "1rem",
        zIndex: "1"
    },
    actionContainer: {
        gridColumnStart: "5",
        gridColumnEnd: "8",
        gridRowStart: "1",
        gridRowEnd: "1",
        zIndex: "2",
        display: "flex",
        alignContent: "space-between",
        justifyContent: "flex-end"
    },
    clickableContainer: {
        cursor: "pointer",
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        marginLeft: "1rem"
    },
    actionText: {
        color: colors.primary01Color100,
        ...typography.actionText,
        display: "inline",
    },
    actionIcon: {
        display: "inline"
    },
    inputStyle: {
        "&.MuiInput-underline:before": {
            borderBottomColor: colors.primary01Color100,
        },
        "&.MuiInput-underline.Mui-disabled:before": {
            borderBottomColor: colors.lineColor,
            borderBottomStyle: "solid"
        },
        "&.Mui-disabled": {
            color: colors.text03Color
        },
        color: colors.text01Color
    },
    labelStyle: {
        "&.Mui-disabled": {
            color: colors.text01Color,
        },
        color: colors.text01Color,
        ...typography.actionText,
    },
    savingContainer: {
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
    },
    savingText: {
        marginLeft: "0.5rem"
    },
}));

// The EditorFieldComponent renders the text field for user to edit the locale value.
// It also manages the state of it's parent component to ensure that only 1 locale value is being updated at any point in time.
const EditorFieldComponent = ({localeObject, localeKey, editing, setEditing, widgetData, setWidgetData, initState, saving, typeKey, valueKey, index, updateLocalisationValueFunc}) => {
    const classes = useStyles();
    const localeValueObject = Object.values(localeObject)[0]
    const localeValue = localeValueObject.value
    const localeEditing = localeValueObject.editing
    return (
        <Grid 
            item 
            xs={12} 
            key={localeKey}
            className={clsx(classes.editorFieldContainer)}
        >
            <TextField
                className={clsx(classes.textField)}
                label={localeKey}
                fullWidth
                disabled={!localeEditing}
                value={localeValue}
                InputProps={{
                    className: clsx(classes.inputStyle),
                }}
                InputLabelProps={{
                    className: clsx(classes.labelStyle),
                }}
                onChange={(event) => { 
                    const newWidgetData = {
                        ...widgetData
                    }
                    newWidgetData[typeKey][valueKey].parsedKeyValue[index][localeKey].value = event.target.value
                    setWidgetData(newWidgetData)
                }} 
            />
            <Box
                className={localeEditing || editing ? clsx(classes.hide) : clsx(classes.actionContainer)}
            >
                <Box
                    onClick={() => {
                        const newWidgetData = {
                            ...widgetData
                        }
                        newWidgetData[typeKey][valueKey].parsedKeyValue[index][localeKey].editing = true
                        setWidgetData({
                            ...newWidgetData
                        })
                        setEditing(true)
                    }}
                    className={clsx(classes.clickableContainer)}
                >
                    <img className={clsx(classes.actionIcon)} src={EditIcon} alt="edit"/>
                    <Typography className={clsx(classes.actionText)}>Edit</Typography>
                </Box>
            </Box>
            <Box
                className={saving || !localeEditing ? clsx(classes.hide) : clsx(classes.actionContainer)}
            >
                <Box
                    onClick={() => {
                        if(!saving) {
                            setEditing(false)
                            initState()
                        }
                    }}
                    className={clsx(classes.clickableContainer)}
                >
                    <img className={clsx(classes.actionIcon)} src={DiscardIcon} alt="discard"/>
                    <Typography className={clsx(classes.actionText)}>Discard</Typography>
                </Box>
                <Box
                    onClick={() => {
                        if (!saving) {
                            updateLocalisationValueFunc(localeKey, localeValue)
                        }
                    }}
                    className={clsx(classes.clickableContainer)}
                >
                    <img className={clsx(classes.actionIcon)} src={CheckIcon} alt="save"/>
                    <Typography className={clsx(classes.actionText)}>Save</Typography>
                </Box>
            </Box>
            {saving && localeEditing && <Box className={clsx(classes.actionContainer, classes.savingContainer)}>
                <CircularProgress 
                    size="12px"
                    style={{'color': colors.primary01Color100}}
                />
                <Typography className={clsx(classes.actionText, classes.savingText)}>Saving</Typography>
            </Box>}
        </Grid>
    )
}

export default EditorFieldComponent;