import React, { useRef, useEffect, useState } from 'react';
import Chart from "chart.js/auto";
import Box from '@material-ui/core/Box';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const chartConfig = {
    type: "doughnut",
    options: {
      cutout: "90%",
    },
    data: {
      datasets: [{
        data: [0, 1],
        backgroundColor: [
            colors.secondary02Color100
        ],
        hoverOffset: 4
      }]
    }
  };

const useStyles = makeStyles(() => ({
    widgetContainer: {
        marginTop: "1rem",
        marginBottom: "1rem",
    },
    chartContainer: {
        display: "grid",
        gridTemplateColumns: "repeat(10, 1fr)",
        gridTemplateRows: "repeat(7, 2rem)",
        gridColumnGap: "0px",
        gridRowGap: "0px",
    },
    planDetails: {
        gridArea: "1 / 8 / 1 / 11",
        backgroundColor: colors.secondary04Color100,
        textAlign: "center",
        padding: "5px 10px 5px 10px",
        margin: "auto 0px",
        borderTopLeftRadius: "15px",
        borderBottomLeftRadius: "15px",
        color: colors.general01Color,
        ...typography.heading3
    },
    speed: {
        fontSize: "2rem",
        fontWeight: "800",
        gridArea: "3 / 4 / 4 / 8",
        marginTop: "1rem",
        textAlign: "center"
    },
    label: {
        ...typography.body2,
        textAlign: "right",
        gridArea: "4 / 4 / 5 / 7",
        marginTop: "1.3rem",
    },
    image: {
        maxWidth: "1rem",
        marginLeft: "0.2rem",
        marginTop: "1.3rem",
        gridArea: "4 / 7 / 5 / 8"
    },
    chart: {
        gridArea: "1 / 3 / 8 / 9"
    },
}));

const UserBaseCard = ({widget, updateEditorComponent, resetEditorComponent, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const chartId = "doughnut"
    const chartContainer = useRef(chartId);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        if (shouldRenderEditorComponent && widgetData.length != 0) {
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
                    resetFunction={resetEditorComponent}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage])

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget) return
        setKeyValueToState(widget, "title", "components[0].data.gauge.title", setWidgetData)
        setKeyValueToState(widget, "description", "components[0].data.gauge.description", setWidgetData)
        setKeyValueToState(widget, "tooltip.image", "components[0].data.gauge.tooltip.image", setWidgetData)
    }

    useEffect(() => {
        initState()
    }, [widget])
    
    useEffect(() => {
        if (chartContainer && chartContainer.current) {
            // This throws a lint error in sonarqube but its unavoidalble as the new Chart invocation is needed to instantiate the chart as the chart is being rendered using a ref
            new Chart(chartContainer.current, chartConfig);
        }
    }, [chartContainer]);
    
    const dataLeft = getValueFromWidgetData(widgetData, "title", "components[0].data.gauge.title")
    const dataLeftDescription = getValueFromWidgetData(widgetData, "description", "components[0].data.gauge.description")
    const dataLeftToolTip = getValueFromWidgetData(widgetData, "tooltip.image", "components[0].data.gauge.tooltip.image")
    return (<Box className={clsx(classes.widgetContainer)}>
        <Box className={clsx(classes.chartContainer)}>
            <Box className={clsx(classes.planDetails)}>
                Plan Details
            </Box>
            <span className={clsx(classes.speed)}>{dataLeft}</span>
            <span className={clsx(classes.label)}>{dataLeftDescription}</span>
            <img src={dataLeftToolTip} alt="tooltip" className={clsx(classes.image)} />
            <Box className={clsx(classes.chart)}>
                <canvas 
                    id={chartId} 
                    ref={chartContainer}
                />
            </Box>
        </Box>
    </Box>)
}

export default UserBaseCard;