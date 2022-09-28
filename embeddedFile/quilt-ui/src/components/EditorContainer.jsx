import React from 'react';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';

import PrimaryButton from './PrimaryButton';
import { colors, typography } from '../styles/globalStyles';

const useStyles = makeStyles(() => ({
    previewWarning: {
        backgroundColor: colors.previewWarningBackgroundColor,
        padding: "0.5rem",
        marginBottom: "1rem",
        ...typography.warningText,
        color: colors.previewWarningColor
    },
    editorCard: {
        padding: "2rem"
    },
    buttonContainer: {
        marginTop: "1rem",
        float: "right"
    },
    errorText: {
        color: colors.alert01Color100
    }
}));

// EditorContainer renders the container containing a collection of EditorFieldComponents
// The clickableCards prop renders cards that lets a user click into nested EditorContainers. This is for complex widgets that contains nested components.
const EditorContainer = ({errorMessage, editorFieldComponents, resetFunction, clickableCards}) => {
    const classes = useStyles();
    
    return (<Box>
        <Card
            className={clsx(classes.previewWarning)}
        >
            <Typography>Please note that changes shown here is only a preview and should still be tested on an actual  device</Typography>
        </Card>
        <Card
            className={clsx(classes.editorCard)} 
        >
            <Grid container direction="row">
                {editorFieldComponents}
                {clickableCards}
                {errorMessage && <Grid item xs={12}>
                    <Typography className={clsx(classes.errorText)} variant="body1">Error: {errorMessage}</Typography>
                </Grid>}
                <Grid item xs={12}>
                    <Box
                        className={clsx(classes.buttonContainer)}
                    >
                        <PrimaryButton
                            onClick={()=>{
                                resetFunction()
                            }}
                            buttonText="Done"
                        />
                    </Box>
                </Grid>
            </Grid>
        </Card>
    </Box>)
}

export default EditorContainer;