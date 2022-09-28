import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import LinearProgress from '@material-ui/core/LinearProgress';
import clsx from 'clsx';
import { makeStyles, withStyles } from '@material-ui/core/styles';

import { map } from 'lodash';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const ProgressBar = withStyles((theme) => ({
    root: {
      height: 10,
      borderRadius: 5,
    },
    colorPrimary: {
      backgroundColor: colors.background03Color
    },
    bar: {
      borderRadius: 5,
      backgroundColor: colors.secondary01Color100
    },
}))(LinearProgress);

const useStyles = makeStyles(() => ({
    widgetContainer: {
        margin: "1rem",
        border: `2px solid ${colors.primary01Color100}`,
        borderRadius: "0.4rem",
        boxShadow: `5px 5px ${colors.primary01Color100}`
    },
    header: {
        margin: "1rem 0px",
        ...typography.heading3,
        textAlign: "center",
        color: colors.general01Color,
        padding: "0.2rem 0px",
        backgroundImage: colors.gradient01,
        width: "100.5%",
        marginLeft: "-1px"
    },
    tileProgressTrackerContainer: {
        margin: "0.75rem",
    },
    tileImage: {
        maxWidth: "80%"
    },
    rewardImage: {
        float: "right",
        maxWidth: "1rem"
    },
    rewardText: {
        color: colors.text02Color,
        ...typography.body2,
        textAlign: "left"
    },
    tileTitle: {
        ...typography.heading3,
        marginBottom: "0.5rem"
    },
    tileProgressionText: {
        ...typography.body2,
    }
}));

const BorderedListItem = ({widget, updateEditorComponent, resetSelectedComponentCard, shouldRenderComponentEditorComponent, pageId, fetchPageData}) => {
    const [previewComponents, setPreviewComponents] = useState(null);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const classes = useStyles();
    
    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        if (shouldRenderComponentEditorComponent && widgetData) {
            // updateLocalisationValueFunc is a wrapper function for updateLocalisationValue to avoid having to pass in params(that are local to this scope) to renderEditorFieldComponents function.
            const updateLocalisationValueFunc = (localeKey, localeValue) => {
                updateLocalisationValue({localeKey, localeValue, pageId, setErrorMessage, setSaving, setEditing, fetchPageData})
            }
            // Generate EditorFieldComponents based on widgetData
            const editorFieldComponents = renderEditorFieldComponents({widgetData, saving, setWidgetData, editing, setEditing, initState, updateLocalisationValueFunc})
            // Invoke updateEditorComponent to render the EditorComponent
            updateEditorComponent(
                <EditorContainer 
                    errorMessage={errorMessage} 
                    editorFieldComponents={editorFieldComponents}
                    resetFunction={resetSelectedComponentCard}
                />
            )
        }
    }, [shouldRenderComponentEditorComponent, widgetData, saving, errorMessage])

    // Render Preview component
    useEffect(() => {
        setPreviewComponents(
            map(widgetData, (value, type) => {
                if(type === "tile-progress-tracker") {
                    const parsedId = getValueFromWidgetData(widgetData, type, "data.id")
                    const parsedTileImage = getValueFromWidgetData(widgetData, type, "data.tileImage")
                    const parsedTitle = getValueFromWidgetData(widgetData, type, "data.title")
                    const parsedDescription = getValueFromWidgetData(widgetData, type, "data.description")
                    const parsedRewardImage = getValueFromWidgetData(widgetData, type, "data.rewardImage")
                    const parsedRewardText = getValueFromWidgetData(widgetData, type, "data.rewardText")
                    const parsedProgressionText = getValueFromWidgetData(widgetData, type, "data.progressionText")
                    const parsedStatus = getValueFromWidgetData(widgetData, type, "data.status")
                    const parsedRewardType = getValueFromWidgetData(widgetData, type, "data.reward.type")
                    const parsedRewardBenefitTitle = getValueFromWidgetData(widgetData, type, "reward.benefit.title")
                    const parsedRewardBenefitTNC = getValueFromWidgetData(widgetData, type, "reward.benefit.tnc")
                    const [progressNumerator , unusedValue, progressDenominator] = parsedProgressionText.split(" ")
                    const progressValue = Math.round((parseInt(progressNumerator)/parseInt(progressDenominator))*100)
                    return(<Box 
                        className={clsx(classes.tileProgressTrackerContainer)}
                        id={parsedId}
                        key={parsedId}
                    >
                        <Grid container direction="row">
                            <Grid item xs={2}>
                                <img src={parsedTileImage} alt="tile image" className={clsx(classes.tileImage)} />
                            </Grid>
                            <Grid item xs={10}>
                                <Grid container direction="row">
                                    <Grid item xs={8}>
                                        <Typography 
                                            variant="body1"
                                            className={clsx(classes.tileTitle)}
                                        >
                                            {parsedTitle}
                                        </Typography>
                                    </Grid>
                                    <Grid item xs={3}>
                                        <Typography 
                                            variant="body1"
                                            className={clsx(classes.rewardText)}
                                        >
                                            {parsedRewardText}
                                        </Typography>
                                    </Grid>
                                    <Grid item xs={1}>
                                        <img src={parsedRewardImage} alt="reward image" className={clsx(classes.rewardImage)} />
                                    </Grid>
                                    <Grid item xs={12}>
                                        <ProgressBar 
                                            variant="determinate" 
                                            value={progressValue}
                                        />
                                    </Grid>
                                </Grid>
                                <Typography 
                                    variant="body1"
                                    className={clsx(classes.tileProgressionText)}
                                >
                                    {parsedProgressionText}
                                </Typography>
                            </Grid>
                        </Grid>
                    </Box>)
                }
            })
        )
    }, [widgetData])

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget) return
        switch(widget.type) {
            case "tile-progress-tracker" : {
                setKeyValueToState(widget, "tile-progress-tracker", "data.id", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.tileImage", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.title", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.description", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.rewardImage", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.rewardText", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.progressionText", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.status", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.reward.type", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.reward.benefit.title", setWidgetData)
                setKeyValueToState(widget, "tile-progress-tracker", "data.reward.benefit.tnc", setWidgetData)
            }
        }
    }

    useEffect(() => {
        initState()
    }, [widget])

    
    return (previewComponents && <>
        {map(previewComponents, previewComponent => previewComponent)}
    </>)
}

export default BorderedListItem;